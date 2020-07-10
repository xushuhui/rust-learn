package main

import (
	"skr-shop-cms-api/core"
	"skr-shop-cms-api/model"
	"skr-shop-cms-api/route"
)

func main() {
	r := route.InitRouter()
	model.InitDB()
	core.InitRedis()
	core.InitValidate()
	r.Run()
}
