package domain

type Token struct {
	TokenID   string    `json:"token_id"`
	ProjectID string    `json:"project_id"`
	Authority Authority `json:"authority"`
}

type Authority string

const (
	AuthorityOwner Authority = "owner"
	ReadAndWrite   Authority = "read_and_write"
	ReadOnly       Authority = "read_only"
)
