package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/schema-creator/services/sql-service/cmd/app/config"
	"github.com/schema-creator/services/sql-service/internal/project/infrastruture/mongodb"
	"github.com/schema-creator/services/sql-service/internal/project/interfaces/http"
	userhttp "github.com/schema-creator/services/sql-service/internal/user/interfaces/http"
	wshttp "github.com/schema-creator/services/sql-service/internal/websocket/interfaces/http"
	"github.com/schema-creator/services/sql-service/pkg/logger"
	"github.com/schema-creator/services/sql-service/pkg/mongo"
)

// @title go-schema
// @version 1.0
// @description スキーマ設計API
// @contact.name murasame
// @contact.email [fake]
// @host localhost:8080
// @BasePath /v1/
func main() {
	l := logger.NewLogger(logger.DEBUG)
	config.LoadEnv(l)

	ctx := context.Background()

	conn := mongo.NewConnection(l)
	defer conn.Close(ctx)

	db, err := conn.Connection(config.Config.Mongo.Database)
	if err != nil {
		l.Error(err)
		return
	}
	gin := gin.Default()
	router := userhttp.NewUserRouter(gin, l, db)
	http.NewProjectRouter(gin, l, mongodb.NewProjectRepository(db))
	wshttp.NewWsRouter(gin, l, db)
	router.Run(config.Config.Server.Addr)
}
