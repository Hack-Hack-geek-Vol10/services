package domain

type Token struct {
	TokenID   string `json:"token_id"`
	ProjectID string `json:"project_id"`
	Token     string `json:"token"`
	Authority int    `json:"authority"`
}
