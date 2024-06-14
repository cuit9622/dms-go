package api

import (
	"context"
	"github.com/cuit9622/dms/common/global"
	"github.com/cuit9622/dms/common/pb"
	"github.com/cuit9622/dms/common/response"
	"github.com/cuit9622/dms/common/response/errors"
	"github.com/cuit9622/dms/common/util/grpcUtil"
	"github.com/cuit9622/dms/dorm/client"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"strconv"
)

func updateDormBed(c *gin.Context) {
	bed := DormBed{}
	err := c.ShouldBindJSON(&bed)
	if err != nil {
		response.ErrorCode(c, errors.BAD_REQUEST)
		return
	}
	student := Student{}
	if bed.StudentNo != "" {
		global.GLO_HTTP_CLIENT.GetWithPathVariable("student",
			"/stuNum",
			bed.StudentNo,
			&student)
		if student.ID == 0 {
			response.Error(c, 400, "该学号未匹配学生")
			return
		}
	}
	pbBed := pb.DormBed{
		Id:        bed.ID,
		DormID:    bed.DormID,
		StudentID: student.ID,
	}

	var r int32
	grpcUtil.CallGrpc("dorm-service", func(con *grpc.ClientConn, ctx context.Context) error {
		service := client.GetDormBedService(con)
		var err error
		result, err := service.Update(ctx, &pbBed)
		r = result.Value
		return err
	})
	if r == -1 {
		response.Error(c, 500, "已达寝室容量上限")
		return
	}
	response.Success(c, r)
}
func deleteDormBed(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	var r int32
	grpcUtil.CallGrpc("dorm-service", func(con *grpc.ClientConn, ctx context.Context) error {
		service := client.GetDormBedService(con)
		var err error
		result, err := service.Delete(ctx, &wrapperspb.Int64Value{Value: id})
		r = result.Value
		return err
	})
	response.Success(c, r)
}
