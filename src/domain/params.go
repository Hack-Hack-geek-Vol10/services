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

type CreateMemberParam struct {
	ProjectID string    `json:"project_id"`
	UserID    string    `json:"user_id"`
	Authority Authority `json:"authority"`
}

type UpdateAuthorityParam struct {
	ProjectID string    `json:"project_id"`
	UserID    string    `json:"user_id"`
	Authority Authority `json:"authority"`
}

type DeleteMemberParam struct {
	ProjectID string `json:"project_id"`
	UserID    string `json:"user_id"`
}

type ReadProjectsParam struct {
	UserID string `json:"user_id"`
	Limit  uint32 `json:"limit"`
	Offset uint32 `json:"offset"`
}
