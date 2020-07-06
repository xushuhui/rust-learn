package main

import "skr-shop-cms-api/core"

func main() {
	r := core.SetupRouter()
	r.Run()
}
