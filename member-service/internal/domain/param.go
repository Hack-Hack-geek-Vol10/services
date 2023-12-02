package domain

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
