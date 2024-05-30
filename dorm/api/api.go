package api

import (
	"context"
	"fmt"

	"github.com/cuit9622/dms/common/proto"
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

type person struct {
	name string
	age  int8
}

func Test1(c *gin.Context) {
	var r *proto.HelloReply
	grpcUtil.CallGrpc("dorm-service", func(con *grpc.ClientConn, ctx context.Context) error {
		client := client.GetGreeterService(con)
		var err error
		r, err = client.SayHello(ctx, &proto.IdRequest{Id: 1})
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
