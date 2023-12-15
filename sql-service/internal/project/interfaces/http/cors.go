package http

import "github.com/gin-contrib/cors"

func (r *UserRouter) setupCors() {
	r.Gin.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"content-type", "dbauthorization"},
	}))
}
