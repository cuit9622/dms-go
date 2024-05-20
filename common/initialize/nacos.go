package initialize

import (
	"cuit9622/dms-common/global"
	"fmt"
	"net"
	"strconv"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func InitNacos() string {
	server := global.GLO_VP.GetString("NACOS_SERVER")
	serverIP, serverPortStr, err := net.SplitHostPort(server)
	if err != nil {
		panic(fmt.Errorf("fatal error parse nacos server: %s", err))
	}
	serverPort, _ := strconv.ParseUint(serverPortStr, 10, 64)
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: serverIP,
			Port:   serverPort,
		},
	}

	clientConfig := constant.ClientConfig{
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "debug",
	}
	nacosClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(fmt.Errorf("fatal error create nacosClient: %s", err))
	}
	global.GLO_NACOS = &nacosClient
	return serverIP
}
