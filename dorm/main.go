package main

import (
	"cuit9622/dms-common/initialize"
	"dms-dorm/api"
)

func main() {
	g, ln := initialize.InitSecurity()
	api.SetRouter(g)
	initialize.RunServer(g, ln)
}
