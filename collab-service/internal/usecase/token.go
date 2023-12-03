package usecase

import (
	"context"

	"github.com/google/uuid"
	token "github.com/schema-creator/services/collab-service/api/v1"
	"github.com/schema-creator/services/collab-service/internal/domain"
	"github.com/schema-creator/services/collab-service/internal/infra"
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

func (t *tokenService) CreateToken(ctx context.Context, arg *token.CreateTokenRequest) (*token.CreateTokenResponse, error) {
	param := domain.CreateTokenParam{
		TokenID:   uuid.New().String(),
		ProjectID: arg.ProjectId,
		Authority: domain.Authority(arg.Authority),
	}

	err := t.tokenRepo.Create(ctx, param)
	if err != nil {
		return nil, err
	}

	return &token.CreateTokenResponse{
		Token: param.TokenID,
	}, nil
}

func (t *tokenService) GetToken(ctx context.Context, arg *token.GetTokenRequest) (*token.GetTokenResponse, error) {
	param := domain.GetTokenParam{
		TokenID: arg.Token,
	}
	tokenInfo, err := t.tokenRepo.Get(ctx, param)
	if err != nil {
		return nil, err
	}

	return &token.GetTokenResponse{
		TokenId:   tokenInfo.TokenID,
		ProjectId: tokenInfo.ProjectID,
		Authority: string(tokenInfo.Authority),
	}, nil
}

func (t *tokenService) DeleteToken(ctx context.Context, arg *token.DeleteTokenRequest) (*token.DeleteTokenResponse, error) {
	param := domain.DeleteTokenParam{
		ProjectID: arg.ProjectId,
	}
	err := t.tokenRepo.Delete(ctx, param)
	if err != nil {
		return nil, err
	}

	return &token.DeleteTokenResponse{
		ProjectId: arg.ProjectId,
	}, nil
}
