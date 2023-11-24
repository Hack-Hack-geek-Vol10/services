package domain

type Project struct {
	ProjectID  string `json:"project_id"`
	Title      string `json:"title"`
	LastImage  string `json:"last_image"`
	IsPersonal bool   `json:"is_personal"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	IsDelete   bool   `json:"is_delete"`
}

type User struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}

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
