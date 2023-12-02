package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/schema-creator/services/project-service/internal/domain"
	token "github.com/schema-creator/services/token-service/api/v1"
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

	result, err := t.tokenRepo.Create(ctx, param)
	if err != nil {
		return nil, err
	}

	return &token.TokenDetails{}
}
