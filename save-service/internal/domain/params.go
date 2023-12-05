package domain

type CreateSaveParam struct {
	SaveID    string `json:"save_id"`
	ProjectID string `json:"project_id"`
	Editor    string `json:"editor"`
	Object    string `json:"object"`
	CreatedAt string `json:"created_at"`
}

type GetSaveParam struct {
	ProjectID string `json:"project_id"`
}

type DeleteSaveParam struct {
	ProjectID string `json:"project_id"`
}
