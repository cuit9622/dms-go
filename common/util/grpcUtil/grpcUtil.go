package grpcUtil

import (
	"context"
	"fmt"
	"time"

	"github.com/cuit9622/dms/common/global"
	"github.com/cuit9622/dms/common/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CallGrpc(serverName string, body func(*grpc.ClientConn, context.Context) error) {
	ins := util.GetInstance(serverName)
	addr := fmt.Sprintf("%s:%d", ins.Ip, ins.Port)
	con, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		global.GLO_LOG.Panic(fmt.Sprintf("failed connect to %s\n%s", addr, err.Error()))
	}
	defer con.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = body(con, ctx)
	if err != nil {
		global.GLO_LOG.Panic(fmt.Sprintf("failed call grpc: %s", err.Error()))
	}
}
