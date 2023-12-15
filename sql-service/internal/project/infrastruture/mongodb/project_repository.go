package mongodb

import (
	"context"

	"github.com/schema-creator/services/sql-service/internal/project/domain/repository"
	"github.com/schema-creator/services/sql-service/internal/project/domain/values"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ProjectCollection = "projects"
)

type projectRepository struct {
	db *mongo.Database
}

func NewProjectRepository(db *mongo.Database) repository.ProjectRepository {
	return &projectRepository{db}
}

func (pr *projectRepository) Create(ctx context.Context, args values.CreateProject) (*values.Project, error) {
	_, err := pr.db.Collection(ProjectCollection).InsertOne(ctx, args)
	return &values.Project{
		ProjectID: args.ProjectID,
		Title:     args.Title,
		OwnerID:   args.OwnerID,
		Users:     args.Users,
		CreateAt:  args.CreatedAt,
		UpdateAt:  args.UpdateAt,
	}, err
}

func (pr *projectRepository) GetProjectsByUserID(ctx context.Context, userID string) ([]*values.Project, error) {
	var projects []*values.Project
	cursor, err := pr.db.Collection(ProjectCollection).Find(ctx, bson.M{"owner_id": userID})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

func (pr *projectRepository) DeleteAllByUserID(ctx context.Context, userID string) error {
	_, err := pr.db.Collection(ProjectCollection).DeleteMany(ctx, bson.M{"owner_id": userID})
	return err
}

func (pr *projectRepository) GetOneByID(ctx context.Context, id string) (*values.Project, error) {
	var project values.Project
	if err := pr.db.Collection(ProjectCollection).FindOne(ctx, bson.M{"_id": id}).Decode(&project); err != nil {
		return nil, err
	}
	return &project, nil
}

func (pr *projectRepository) UpdateProject(ctx context.Context, args values.Project) (*values.Project, error) {
	if _, err := pr.db.Collection(ProjectCollection).UpdateOne(ctx, bson.M{"_id": args.ProjectID}, bson.M{"$set": args}); err != nil {
		return nil, err
	}
	return pr.GetOneByID(ctx, args.ProjectID)
}

func (pr *projectRepository) Delete(ctx context.Context, id string) error {
	return pr.db.Collection(ProjectCollection).FindOneAndDelete(ctx, bson.M{"_id": id}).Err()
}

func (pr *projectRepository) JoinProject(ctx context.Context, projectID, userID string) (*values.Project, error) {
	var project *values.Project
	if _, err := pr.db.Collection(ProjectCollection).UpdateOne(ctx, bson.M{"_id": projectID}, bson.M{"$push": bson.M{"users": userID}}); err != nil {
		return nil, err
	}

	if err := pr.db.Collection(ProjectCollection).FindOne(ctx, bson.M{"_id": projectID}).Decode(&project); err != nil {
		return nil, err
	}
	return project, nil
}

func (pr *projectRepository) WsUpdateEditor(ctx context.Context, projectID, editor string) (*values.Project, error) {
	var project *values.Project
	if _, err := pr.db.Collection(ProjectCollection).UpdateOne(ctx, bson.M{"_id": projectID}, bson.M{"$set": bson.M{"editor": editor}}); err != nil {
		return nil, err
	}
	return project, nil
}

func (pr *projectRepository) WsUpdateObject(ctx context.Context, projectID, object string) (*values.Project, error) {
	var project *values.Project
	if _, err := pr.db.Collection(ProjectCollection).UpdateOne(ctx, bson.M{"_id": projectID}, bson.M{"$set": bson.M{"object": object}}); err != nil {
		return nil, err
	}
	return project, nil
}

func (pr *projectRepository) GetEditorData(ctx context.Context, projectID string) (string, error) {
	var project *values.Project
	if err := pr.db.Collection(ProjectCollection).FindOne(ctx, bson.M{"_id": projectID}).Decode(&project); err != nil {
		return "", err
	}
	return project.Editor, nil
}
