package client

import (
	"github.com/cuit9622/dms/common/pb"
	"google.golang.org/grpc"
)

func GetDormBuildingService(con *grpc.ClientConn) pb.DormBuildingServiceClient {
	return pb.NewDormBuildingServiceClient(con)
}
