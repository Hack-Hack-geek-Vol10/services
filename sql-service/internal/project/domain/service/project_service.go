package service

import (
	"context"

	"github.com/schema-creator/services/sql-service/internal/project/domain/response"
	"github.com/schema-creator/services/sql-service/pkg/ddl"
)

type ProjectService interface {
	ConvertDDL(ctx context.Context, id string, convertType ddl.ConvertType) (*response.ConvertDDL, error)
}
