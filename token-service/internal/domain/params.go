package domain

type CreateTokenParam struct {
	TokenID   string    `json:"token_id"`
	ProjectID string    `json:"project_id"`
	Authority Authority `json:"authority"`
}

type GetTokenParam struct {
	TokenID string `json:"token_id"`
}

type DeleteTokenParam struct {
	ProjectID string `json:"project_id"`
}
