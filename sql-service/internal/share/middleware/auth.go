package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/schema-creator/services/sql-service/pkg/jwt"
	"github.com/schema-creator/services/sql-service/pkg/logger"
)

const (
	AuthorizationHeaderKey = "dbauthorization"
	AuthorizationType      = "dbauthorization_type"
	AuthorizationClaimsKey = "authorization_claim"
)

type Middleware struct {
	l logger.Logger
}

func NewMiddleware(l logger.Logger) *Middleware {
	return &Middleware{l}
}

func (m *Middleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)
		// if len(authorizationHeader) == 0 {
		// 	m.l.Info("authorization header is not provided")
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
		// 	return
		// }

		// m.decodeJwt(ctx, authorizationHeader)
	}
}

func (m *Middleware) decodeJwt(ctx *gin.Context, header string) {
	fields := strings.Fields(header)
	if len(fields) < 1 {
		m.l.Info("invalid authorization header format")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
		return
	}

	payload, err := jwt.ValidJWT(fields[0])
	if err != nil {
		m.l.Infof("invalid access token : %v", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
	}
	ctx.Set(AuthorizationClaimsKey, payload)
	ctx.Next()
}
