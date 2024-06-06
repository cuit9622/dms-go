package main

import (
	"github.com/cuit9622/dms/common/initialize"
	"github.com/cuit9622/dms/common/pb"
	"github.com/cuit9622/dms/dorm-service/service"
	"google.golang.org/grpc"
)

func main() {
	ln := initialize.InitCommon()
	initialize.InitGorm()

	server := grpc.NewServer()
	pb.RegisterDormBuildingServiceServer(server, service.DormBuildingService{})

	initialize.RunGrpcServer(server, ln)
}
