package route

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"skrshop-cms-api/core"
	"time"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	gin.SetMode(core.RunMode)
	r := InitRouter()
	port := core.HTTPPort
	HttpSrvHandler = &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  core.ReadTimeout,
		WriteTimeout: core.WriteTimeout,
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", port)
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", port, err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
