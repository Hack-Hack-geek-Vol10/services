package usecase

import (
	"context"

	"github.com/google/uuid"
	token "github.com/schema-creator/services/token-service/api/v1"
	"github.com/schema-creator/services/token-service/internal/domain"
	"github.com/schema-creator/services/token-service/internal/infra"
)

type tokenService struct {
	token.UnimplementedTokenServiceServer
	tokenRepo infra.TokenRepo
}

func NewTokenService(tokenRepo infra.TokenRepo) token.TokenServiceServer {
	return &tokenService{
		tokenRepo: tokenRepo,
	}
}

func (t *tokenService) CreateToken(ctx context.Context, arg *token.CreateProjectRequest) {
	param := domain.CreateTokenParam{
		TokenID:   uuid.New().String(),
		ProjectID: arg.ProjectId,
		Authority: arg.Authority,
	}

	_, err := t.tokenRepo.Create(ctx, param)
	if err != nil {
		return nil, err
	}

	return &token.CreateTokenResponse{
		Token: param.TokenID,
	}, nil
}

func (t *tokenService) GetToken(ctx context.Context, arg *token.GetTokenRequest) (*token.GetTokenResponse, error) {
	tokenInfo, err := t.tokenRepo.ReadOne(ctx, arg.Token)
	if err != nil {
		return nil, err
	}

	return &token.GetTokenResponse{
		Token: tokenInfo.TokenID,
	}, nil
}
