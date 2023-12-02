package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	dbconn "github.com/murasame29/db-conn/sqldb/postgres"
	"github.com/schema-creator/services/migrate-service/cmd/config"
	"github.com/schema-creator/services/migrate-service/src/server"
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

	router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if token != config.Config.Server.Token {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "unauthorized",
			})
		}
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	router.POST("/migrate", func(c *gin.Context) {
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			c.AbortWithError(500, err)
		}
		m, err := migrate.NewWithDatabaseInstance(
			"file://app/migrations",
			"postgres", driver)
		if err != nil {
			c.AbortWithError(500, err)
		}
		if err := m.Up(); err != nil {
			c.AbortWithError(500, err)
		}
		c.JSON(200, "success migrate")
	})

	server.Run(router)
}
