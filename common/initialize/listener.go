package initialize

import (
	"fmt"
	"net"
	"strconv"

	"github.com/cuit9622/dms/common/global"
)

func initListener() net.Listener {
	ln, err := net.Listen("tcp", ":"+global.GLO_VP.GetString("server.port"))
	if err != nil {
		panic(fmt.Errorf("fatal error listen port: %s", err))
	}
	_, portStr, _ := net.SplitHostPort(ln.Addr().String())
	port, _ := strconv.ParseUint(portStr, 10, 64)
	global.GLO_INFO = global.Info{
		IP:   findServerIP(),
		Port: port,
		Name: global.GLO_VP.GetString("application.name"),
	}
	global.GLO_LOG.Info(fmt.Sprintf("port binding on: %d", port))
	return ln
}

func findServerIP() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(fmt.Errorf("fatal error get ifaces: %s", err))
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			panic(fmt.Errorf("fatal error get address: %s", err))
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip := v.IP
				if ip.IsPrivate() && !ip.IsLoopback() {
					return ip.String()
				}
			}
		}
	}
	panic(fmt.Errorf("fatal error findIP: %s", err))
}
