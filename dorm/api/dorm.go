package api

import (
	"context"
	"github.com/cuit9622/dms/common/entity"
	"github.com/cuit9622/dms/common/global"
	"github.com/cuit9622/dms/common/model"
	"github.com/cuit9622/dms/common/pb"
	"github.com/cuit9622/dms/common/response"
	"github.com/cuit9622/dms/common/response/errors"
	"github.com/cuit9622/dms/common/util/grpcUtil"
	"github.com/cuit9622/dms/dorm/client"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"strconv"
	"sync"
)

//go:generate gomodifytags -all -add-tags form,json -add-options json=omitempty -transform camelcase --skip-unexported -w -file $GOFILE

type dormGetRequest struct {
	DormBuildingID int64 `form:"dormBuildingID" json:"dormBuildingID,omitempty"`
	Floor          int32 `form:"floor" json:"floor,omitempty"`
	Page           int32 `form:"page" json:"page,omitempty"`
	PageSize       int32 `form:"pageSize" json:"pageSize,omitempty"`
}
type Dorm struct {
	entity.Dorm
	DormBeds []DormBed `form:"dormBeds" json:"dormBeds,omitempty"`
}
type DormBed struct {
	entity.DormBed
	StudentName string `form:"studentName" json:"studentName,omitempty"`
	StudentNo   string `form:"studentNo" json:"studentNo,omitempty"`
}
type Student struct {
	ID     int64  `json:"stuId,omitempty" form:"stuId"`
	Name   string `json:"name,omitempty" form:"name"`
	Sex    int64  `json:"sex" form:"sex"`
	StuNum string `json:"stuNum,omitempty" form:"stuNum"`
}

func getDorms(c *gin.Context) {
	request := dormGetRequest{}
	err := c.ShouldBindQuery(&request)
	if err != nil {
		response.ErrorCode(c, errors.BAD_REQUEST)
		return
	}
	var r *pb.PageResult
	grpcUtil.CallGrpc("dorm-service", func(con *grpc.ClientConn, ctx context.Context) error {
		service := client.GetDormService(con)
		var err error
		r, err = service.Get(ctx, &pb.DormGetRequest{
			DormBuildingID: request.DormBuildingID,
			Floor:          request.Floor,
			Page: &pb.PageRequest{
				Page:     request.Page,
				PageSize: request.PageSize,
			},
		})
		return err
	})
	pageResult := model.PageResult{
		Total: r.Total,
	}
	dst := pb.Dorms{}
	err = anypb.UnmarshalTo(r.Records, &dst, proto.UnmarshalOptions{})
	if err != nil {
		return
	}
	if dst.Dorms == nil {
		dst.Dorms = []*pb.Dorm{}
	}

	wg := sync.WaitGroup{}
	records := make([]Dorm, len(dst.Dorms))
	for i, dorm := range dst.Dorms {
		record := &records[i]
		record.ID = dorm.Id
		record.Name = dorm.Name
		record.Size = dorm.Size
		record.Floor = dorm.Floor

		item := dorm
		wg.Add(1)
		go func() {
			defer wg.Done()

			var lock sync.Mutex
			for _, bed := range item.DormBeds {
				wg.Add(1)
				bed := bed
				go func() {
					defer wg.Done()
					student := Student{}
					err := global.GLO_HTTP_CLIENT.GetWithPathVariable(
						"student",
						"/student",
						strconv.FormatInt(bed.StudentID, 10),
						&student)
					if err != nil {
						global.GLO_LOG.Error(err.Error())
						return
					}
					lock.Lock()
					record.DormBeds = append(record.DormBeds, DormBed{
						DormBed: entity.DormBed{
							ID:        bed.Id,
							StudentID: bed.StudentID,
						},
						StudentName: student.Name,
						StudentNo:   student.StuNum,
					})
					lock.Unlock()
				}()
			}
		}()
	}
	wg.Wait()
	pageResult.Records = &records
	response.Success(c, pageResult)
}
func updateDorm(c *gin.Context) {
	dorm := pb.Dorm{}
	err := c.ShouldBindJSON(&dorm)
	if err != nil {
		response.ErrorCode(c, errors.BAD_REQUEST)
		return
	}
	var r int32
	grpcUtil.CallGrpc("dorm-service", func(con *grpc.ClientConn, ctx context.Context) error {
		service := client.GetDormService(con)
		var err error
		result, err := service.Update(ctx, &dorm)
		r = result.Value
		return err
	})
	response.Success(c, r)
}
func deleteDorm(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	var r int32
	grpcUtil.CallGrpc("dorm-service", func(con *grpc.ClientConn, ctx context.Context) error {
		service := client.GetDormService(con)
		var err error
		result, err := service.Delete(ctx, &wrapperspb.Int64Value{Value: id})
		r = result.Value
		return err
	})
	response.Success(c, r)
}
