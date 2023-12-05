package domain

type CreateEditorParam struct {
	ProjectID string `json:"project_id"`
	Query     string `json:"query"`
}

type GetEditorParam struct {
	ProjectID string `json:"project_id"`
}

type UpdateEditorParam struct {
	ProjectID string `json:"project_id"`
	Query     string `json:"query"`
}

type DeleteEditorParam struct {
	ProjectID string `json:"project_id"`
}
