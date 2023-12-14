package usecase

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/newrelic/go-agent/v3/newrelic"
	token "github.com/schema-creator/services/token-service/api/v1"
	"github.com/schema-creator/services/token-service/internal/domain"
	"github.com/schema-creator/services/token-service/internal/infra"
)

type tokenService struct {
	token.UnimplementedTokenServer
	tokenRepo infra.TokenRepo
}

func NewTokenService(tokenRepo infra.TokenRepo) token.TokenServer {
	return &tokenService{
		tokenRepo: tokenRepo,
	}
}

func (t *tokenService) CreateToken(ctx context.Context, arg *token.CreateTokenRequest) (*token.CreateTokenResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("grpc-CreateToken").End()

	param := domain.CreateTokenParam{
		TokenID:   uuid.New().String(),
		ProjectID: arg.ProjectId,
		Authority: domain.Authority(arg.Authority),
	}

	err := t.tokenRepo.Create(ctx, param)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &token.CreateTokenResponse{
		Token: param.TokenID,
	}, nil
}

func (t *tokenService) GetToken(ctx context.Context, arg *token.GetTokenRequest) (*token.GetTokenResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("grpc-GetToken").End()

	param := domain.GetTokenParam{
		TokenID: arg.Token,
	}
	tokenInfo, err := t.tokenRepo.Get(ctx, param)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &token.GetTokenResponse{
		TokenId:   tokenInfo.TokenID,
		ProjectId: tokenInfo.ProjectID,
		Authority: string(tokenInfo.Authority),
	}, nil
}

func (t *tokenService) DeleteToken(ctx context.Context, arg *token.DeleteTokenRequest) (*token.DeleteTokenResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("grpc-DeleteToken").End()

	param := domain.DeleteTokenParam{
		ProjectID: arg.ProjectId,
	}
	err := t.tokenRepo.Delete(ctx, param)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &token.DeleteTokenResponse{
		ProjectId: arg.ProjectId,
	}, nil
}
