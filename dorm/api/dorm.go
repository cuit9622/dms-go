package api

import (
	"context"
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
)

//go:generate gomodifytags -all -add-tags form -transform camelcase --skip-unexported -w -file $GOFILE

type dormGetRequest struct {
	DormBuildingID int64 `form:"dormBuildingID"`
	Floor          int32 `form:"floor"`
	Page           int32 `form:"page"`
	PageSize       int32 `form:"pageSize"`
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
	pageResult.Records = &dst.Dorms
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
