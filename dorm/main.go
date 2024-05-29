package main

import (
	"github.com/cuit9622/dms/dorm/api"

	"github.com/cuit9622/dms/common/initialize"
)

func main() {
	g, ln := initialize.InitSecurity()
	api.SetRouter(g)
	initialize.RunHttpServer(g, ln)
}
