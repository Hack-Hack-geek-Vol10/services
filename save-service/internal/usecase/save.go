package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	save "github.com/schema-creator/services/save-service/api/v1"
	"github.com/schema-creator/services/save-service/internal/domain"
	"github.com/schema-creator/services/save-service/internal/infra"
)

type saveService struct {
	save.UnimplementedSaveServiceServer
	saveRepo infra.SaveRepo
}

func NewSaveService(saveRepo infra.SaveRepo) save.SaveServiceServer {
	return &saveService{
		saveRepo: saveRepo,
	}
}

func (t *saveService) CreateSave(ctx context.Context, arg *save.CreateSaveRequest) (*save.CreateSaveResponse, error) {
	param := domain.CreateSaveParam{
		SaveID:    uuid.New().String(),
		ProjectID: arg.ProjectId,
		Editor:    arg.Editor,
		Object:    arg.Object,
		CreatedAt: time.Now(),
	}

	err := t.saveRepo.Create(ctx, param)
	if err != nil {
		return nil, err
	}

	return &save.CreateSaveResponse{
		SaveId: param.SaveID,
	}, nil
}

func (t *saveService) GetSave(ctx context.Context, arg *save.GetSaveRequest) (*save.GetSaveResponse, error) {
	param := domain.GetSaveParam{
		ProjectID: arg.ProjectId,
	}
	saveInfo, err := t.saveRepo.Get(ctx, param)
	if err != nil {
		return nil, err
	}

	return &save.GetSaveResponse{
		SaveId: saveInfo.SaveID,
		Editor: saveInfo.Editor,
		Object: saveInfo.Object,
	}, nil
}

func (t *saveService) DeleteSave(ctx context.Context, arg *save.DeleteSaveRequest) (*save.DeleteSaveResponse, error) {
	param := domain.DeleteSaveParam{
		ProjectID: arg.ProjectId,
	}
	err := t.saveRepo.Delete(ctx, param)
	if err != nil {
		return nil, err
	}

	return &save.DeleteSaveResponse{
		ProjectId: arg.ProjectId,
	}, nil
}
