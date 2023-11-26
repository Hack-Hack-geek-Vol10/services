package storages

import (
	"context"
	"database/sql"

	"github.com/Hack-Hack-geek-Vol10/services/src/domain"
)

type memberRepo struct {
	db *sql.DB
}

type MemberRepo interface {
	Create(ctx context.Context, arg domain.CreateMemberParam) error
	ReadAll(ctx context.Context, projectID string) ([]*domain.ProjectMember, error)
}

func NewMemberRepo(db *sql.DB) MemberRepo {
	return &memberRepo{
		db: db,
	}
}

func (m *memberRepo) Create(ctx context.Context, arg domain.CreateMemberParam) error {
	const query = `INSERT INTO project_members (project_id,user_id,authority) VALUES ($1,$2,$3)`
	row := m.db.QueryRowContext(ctx, query, arg.ProjectID, arg.UserID, arg.Authority)
	return row.Err()
}

func (m *memberRepo) ReadAll(ctx context.Context, projectID string) ([]*domain.ProjectMember, error) {
	const query = `SELECT project_id,user_id,authority FROM project_members WHERE project_id = $1`
	row, err := m.db.QueryContext(ctx, query, projectID)
	if err != nil {
		return nil, err
	}
	var projectMembers []*domain.ProjectMember
	for row.Next() {
		var projectMember domain.ProjectMember
		if err := row.Scan(&projectMember.ProjectID, &projectMember.UserID, &projectMember.Authority); err != nil {
			return nil, err
		}
		projectMembers = append(projectMembers, &projectMember)
	}
	return projectMembers, nil
}
