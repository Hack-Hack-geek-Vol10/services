package domain

import "time"

type Save struct {
	SaveID    string    `json:"save_id"`
	ProjectID string    `json:"project_id"`
	Editor    string    `json:"editor"`
	Object    string    `json:"object"`
	CreatedAt time.Time `json:"created_at"`
}
