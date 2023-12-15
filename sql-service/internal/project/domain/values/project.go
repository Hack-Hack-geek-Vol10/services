package values

import "time"

type Project struct {
	ProjectID string    `json:"project_id" bson:"_id"`
	Title     string    `json:"title" bson:"title"`
	OwnerID   string    `json:"owner_id" bson:"owner_id"`
	Users     []string  `json:"users" bson:"users"`
	Object    string    `json:"object" bson:"object"`
	Editor    string    `json:"editor" bson:"editor"`
	CreateAt  time.Time `json:"create_at" bson:"create_at"`
	UpdateAt  time.Time `json:"update_at" bson:"update_at"`
}

type CreateProject struct {
	ProjectID string    `bson:"_id"`
	Title     string    `bson:"title"`
	Users     []string  `bson:"users"`
	OwnerID   string    `bson:"owner_id"`
	CreatedAt time.Time `bson:"created_at"`
	UpdateAt  time.Time `bson:"update_at"`
}
