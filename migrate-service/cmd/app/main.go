package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo/v4"
	dbconn "github.com/murasame29/db-conn/sqldb/postgres"
	"github.com/schema-creator/services/migrate-service/cmd/config"
)

func init() {
	config.LoadEnv()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn := dbconn.NewConnection(
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.DBName,
		config.Config.Database.SSLMode,
		config.Config.Database.ConnectAttempts,
		config.Config.Database.ConnectWaitTime,
		config.Config.Database.ConnectBlocks,
	)
	defer conn.Close(ctx)

	db, err := conn.Connection()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.POST("/migrate", func(c echo.Context) error {
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return err
		}
		m, err := migrate.NewWithDatabaseInstance(
			"file://migrations",
			"postgres", driver)
		if err := m.Up(); err != nil {
			return err
		}
		c.JSON(200, "success migrate")
		return nil
	})

	srv := &http.Server{
		Addr:    ":" + config.Config.Server.ServerAddr,
		Handler: e,
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

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
