package domain

type Save struct {
	SaveID    string `json:"save_id"`
	ProjectID string `json:"project_id"`
	Editor    string `json:"editor"`
	Object    string `json:"object"`
	CreatedAt string `json:"created_at"`
}
