package main

import (
	"cuit9622/dms-common/initialize"
	"dms-dorm/api"
)

func main() {
	g := initialize.InitSecurity()
	api.SetRouter(g)
	g.Run(":8080")
}
