package client

import (
	"github.com/cuit9622/dms/common/proto"
	"google.golang.org/grpc"
)

func GetGreeterService(con *grpc.ClientConn) proto.GreeterClient {
	return proto.NewGreeterClient(con)
}
