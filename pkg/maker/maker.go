package maker

import (
	"time"

	token "github.com/Hack-Hack-geek-Vol10/services/pkg/grpc/token-service/v1"
)

type Maker interface {
	// トークンを作る
	CreateToken(projectID string, authority token.Auth, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
