package request

import "github.com/schema-creator/services/sql-service/pkg/ddl"

type ConvertDDL struct {
	ConvertType ddl.ConvertType `form:"convert_type"`
}
