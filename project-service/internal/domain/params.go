package domain

type CreateProjectParam struct {
	ProjectID string `json:"project_id"`
	Title     string `json:"title"`
}

type ReadProjectsParam struct {
	UserID string `json:"user_id"`
	Limit  uint32 `json:"limit"`
	Offset uint32 `json:"offset"`
}
