package api

import (
	"context"
	"fmt"
	"github.com/cuit9622/dms/common/pb"
	"github.com/cuit9622/dms/common/response/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"strconv"

	"github.com/cuit9622/dms/common/model"
	"github.com/cuit9622/dms/common/response"
	"github.com/cuit9622/dms/common/util/grpcUtil"
	"github.com/cuit9622/dms/common/util/jwtUtil"
	"github.com/cuit9622/dms/dorm/client"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
)

type Test struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func Test1(c *gin.Context) {
	response.ErrorCode(c, errors.FORBIDEN)
}
func createDormBuilding(c *gin.Context) {
	building := pb.DormBuilding{}
	c.ShouldBindJSON(&building)
	var r int32
	grpcUtil.CallGrpc("dorm-service", func(con *grpc.ClientConn, ctx context.Context) error {
		service := client.GetDormBuildingService(con)
		var err error
		result, err := service.Create(ctx, &building)
		r = result.Value
		return err
	})
	response.Success(c, r)
}
func getDormBuildings(c *gin.Context) {
	page := model.PageRequest{}
	err := c.ShouldBindQuery(&page)
	if err != nil {
		return
	}
	var r *pb.PageResult
	grpcUtil.CallGrpc("dorm-service", func(con *grpc.ClientConn, ctx context.Context) error {
		service := client.GetDormBuildingService(con)
		var err error
		r, err = service.Get(ctx, &pb.PageRequest{
			Page:     page.Page,
			PageSize: page.PageSize,
		})
		return err
	})
	pageResult := model.PageResult{
		Total: r.Total,
	}
	dst := pb.DormBuildings{}
	err = anypb.UnmarshalTo(r.Records, &dst, proto.UnmarshalOptions{})
	if err != nil {
		return
	}
	pageResult.Records = &dst.DormBuildings
	response.Success(c, pageResult)
}
func updateDormBuilding(c *gin.Context) {
	building := pb.DormBuilding{}
	err := c.ShouldBindJSON(&building)
	if err != nil {
		return
	}
	var r int32
	grpcUtil.CallGrpc("dorm-service", func(con *grpc.ClientConn, ctx context.Context) error {
		service := client.GetDormBuildingService(con)
		var err error
		result, err := service.Update(ctx, &building)
		r = result.Value
		return err
	})
	response.Success(c, r)
}
func deleteDormBuilding(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	var r int32
	grpcUtil.CallGrpc("dorm-service", func(con *grpc.ClientConn, ctx context.Context) error {
		service := client.GetDormBuildingService(con)
		var err error
		result, err := service.Delete(ctx, &wrapperspb.Int64Value{Value: id})
		r = result.Value
		return err
	})
	response.Success(c, r)
}
func Test2(c *gin.Context) {
	test := &Test{}
	c.ShouldBindBodyWithJSON(test)
	response.Success(c, test)
}
func Test3(c *gin.Context) {
	tokenStr := c.GetHeader("token")
	id, _ := jwtUtil.GetUserId(tokenStr)
	response.Success(c, fmt.Sprintf("Test2 %d", id))
}
