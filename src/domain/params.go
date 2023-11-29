package domain

type CreateProjectParam struct {
	ProjectID string `json:"project_id"`
	Title     string `json:"title"`
}

type CreateUserParams struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type ReadProjectsParam struct {
	UserID string `json:"user_id"`
	Limit  uint32 `json:"limit"`
	Offset uint32 `json:"offset"`
}

type UploadImageParam struct {
	Key         string `json:"key"`
	ContentType string `json:"content_type"`
	Body        []byte `json:"body"`
}
