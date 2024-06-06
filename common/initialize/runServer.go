package initialize

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cuit9622/dms/common/global"

	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"google.golang.org/grpc"
)

func RunHttpServer(g *gin.Engine, ln net.Listener) {
	server := &http.Server{Handler: g}
	go func() {
		if err := server.Serve(ln); err != nil && err != http.ErrServerClosed {
			global.GLO_LOG.Panic(err.Error())
		}
	}()

	waitCloseSignal()

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		global.GLO_LOG.Error(fmt.Sprintf("failed close server: %s", err.Error()))
	}
	closeConnections()
}

func RunGrpcServer(server *grpc.Server, ln net.Listener) {
	go func() {
		if err := server.Serve(ln); err != nil {
			global.GLO_LOG.Panic(err.Error())
		}
	}()

	waitCloseSignal()

	server.GracefulStop()
	closeConnections()
}

func waitCloseSignal() {
	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func closeConnections() {
	//注销Nacos实例并关闭Nacos连接
	global.GLO_NACOS.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          global.GLO_INFO.IP,
		Port:        global.GLO_INFO.Port,
		ServiceName: global.GLO_INFO.Name,
		Ephemeral:   true,
	})
	global.GLO_NACOS.CloseClient()

	// 关闭数据库连接
	if gorm := global.GLO_DB; gorm != nil {
		if db, err := gorm.DB(); err == nil {
			global.GLO_LOG.Info("close gorm")
			db.Close()
		}
	}

	// 关闭Redis连接
	if redis := global.GLO_REDIS; redis != nil {
		global.GLO_LOG.Info("close redis")
		redis.Close()
	}
}
