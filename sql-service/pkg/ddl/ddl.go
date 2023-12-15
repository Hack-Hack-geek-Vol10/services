package ddl

import (
	ddlerror "github.com/schema-creator/services/sql-service/pkg/ddl/error"
	"github.com/schema-creator/services/sql-service/pkg/ddl/postgres"
)

type ConvertType int

const (
	Postgres ConvertType = iota
	MySQL
	SQLServer
	OracleDB
	OBJ
)

func Convert(data string, convertType ConvertType) (string, error) {
	switch convertType {
	case Postgres:
		return postgres.NewPostgres(data).Convert()
	case MySQL:
		//	mysql(data)
	case SQLServer:
		//	sqlServer(data)
	case OracleDB:
		//	oracleDB(data)
	}
	return "", ddlerror.ErrInvalidConvertType
}
