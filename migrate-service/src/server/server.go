package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/schema-creator/services/migrate-service/cmd/config"
)

func Run(router *gin.Engine) {
	srv := &http.Server{
		Addr:    ":" + config.Config.Server.ServerAddr,
		Handler: router,
	}

	go func() {
		log.Println("server started at", config.Config.Server.ServerAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q, os.Interrupt)
	<-q

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
