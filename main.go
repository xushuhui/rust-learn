package main

import (
	"skr-shop-cms-api/core"
	"skr-shop-cms-api/route"
)

func main() {
	r := route.SetupRouter()
	core.InitValidate()
	r.Run()
}
