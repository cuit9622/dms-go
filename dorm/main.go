package main

import (
	"cuit9622/dms-common/initialize"
	"cuit9622/dms-common/jwtUtil"
)

func main() {
	initialize.Init()
	jwtUtil.GetUserId("")
}
