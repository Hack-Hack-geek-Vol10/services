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
