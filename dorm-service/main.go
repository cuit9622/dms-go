package main

import (
	"context"

	"github.com/cuit9622/dms/common/initialize"
	"github.com/cuit9622/dms/common/proto"

	"google.golang.org/grpc"
)

type GreeterServerImpl struct {
	proto.UnimplementedGreeterServer
}

func (i GreeterServerImpl) SayHello(context context.Context, in *proto.IdRequest) (*proto.HelloReply, error) {
	if in.GetId() == 1 {
		return &proto.HelloReply{Name: "Van", Age: 30}, nil
	}
	return &proto.HelloReply{Name: "Xianbei", Age: 114514}, nil
}

func main() {
	ln := initialize.InitCommon()

	grpcServer := grpc.NewServer()
	proto.RegisterGreeterServer(grpcServer, GreeterServerImpl{})

	initialize.RunGrpcServer(grpcServer, ln)
}
