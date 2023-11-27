package services

import (
	"context"
	"time"

	token "github.com/Hack-Hack-geek-Vol10/services/pkg/grpc/token-service/v1"
	"github.com/Hack-Hack-geek-Vol10/services/pkg/maker"
)

type tokenService struct {
	token.UnimplementedTokenServiceServer
	maker maker.Maker
}

func NewTokenService(maker maker.Maker) token.TokenServiceServer {
	return &tokenService{
		maker: maker,
	}
}

func (t *tokenService) CreateToken(ctx context.Context, arg *token.CreateTokenRequest) (*token.CreateTokenResponse, error) {
	pasetoToken, err := t.maker.CreateToken(arg.ProjectId, token.Auth(arg.Authority.Number()), time.Duration(time.Hour*24*7))
	if err != nil {
		return nil, err
	}

	return &token.CreateTokenResponse{
		Token: pasetoToken,
	}, nil
}
func (t *tokenService) VerifyToken(ctx context.Context, arg *token.VerifyTokenRequest) (*token.VerifyTokenResponse, error) {
	payload, err := t.maker.VerifyToken(arg.Token)
	if err != nil {
		return nil, err
	}

	return &token.VerifyTokenResponse{
		ProjectId: payload.ProjectID,
		Authority: token.Auth(payload.Authority),
	}, nil
}
func (t *tokenService) mustEmbedUnimplementedTokenServiceServer() {}
