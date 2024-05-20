package main

import (
	"cuit9622/dms-common/initialize"
	"dms-dorm/api"
	"net/http"
)

func main() {
	g, ln := initialize.InitSecurity()
	api.SetRouter(g)
	http.Serve(ln, g)
}
