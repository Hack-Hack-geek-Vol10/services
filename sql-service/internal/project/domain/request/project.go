package request

import "github.com/schema-creator/services/sql-service/pkg/ddl"

type UserRequestWildcard struct {
	UserID string `uri:"user_id"`
}
type ProjectUserRequestWildcard struct {
	ProjectID string `uri:"project_id"`
	UserID    string `uri:"user_id"`
}

type CreateProject struct {
	Title string `json:"title"`
}

type UpdateProject struct {
	Title string   `json:"title"`
	Users []string `json:"users"`
}

type WsUpdateProject struct {
	Object string `json:"object"`
	Editor string `json:"editor"`
}

type ConvertDDL struct {
	ConvertType ddl.ConvertType `form:"convert_type"`
}
