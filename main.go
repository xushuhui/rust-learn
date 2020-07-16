package main

import (
	"os"
	"os/signal"
	"skrshop-cms-api/core"
	"skrshop-cms-api/route"
	"syscall"
)

func main() {

	core.StartModule()

	route.HttpServerRun()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	route.HttpServerStop()
}
