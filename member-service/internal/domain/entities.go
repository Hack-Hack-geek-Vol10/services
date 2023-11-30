package domain

type ProjectMember struct {
	ProjectID string    `json:"project_id"`
	UserID    string    `json:"user_id"`
	Authority Authority `json:"authority"`
}

type Authority string

const (
	AuthorityOwner Authority = "owner"
	ReadAndWrite   Authority = "read_and_write"
	ReadOnly       Authority = "read_only"
)
