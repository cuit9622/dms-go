package initialize

import (
	"cuit9622/dms-common/global"
	"fmt"
	"net"
	"strconv"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func InitNacos() {
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
	_, err = nacosClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          global.GLO_INFO.IP,
		Port:        global.GLO_INFO.Port,
		ServiceName: global.GLO_INFO.Name,
		Weight:      1,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    global.GLO_VP.GetStringMapString("nacos.metadata"),
	})
	if err != nil {
		panic(fmt.Errorf("fatal error register instance: %s", err.Error()))
	}
	global.GLO_NACOS = &nacosClient
	global.GLO_LOG.Info("Nacos initialization complete")
}
