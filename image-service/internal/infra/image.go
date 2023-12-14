package infra

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/schema-creator/services/image-service/internal/domain"
)

type imageRepo struct {
	app *firebase.App
}

type ImageRepo interface {
	UploadImage(ctx context.Context, arg *domain.UploadImageParam) (string, string, error)
	DeleteImage(ctx context.Context, key string) error
}

func NewImageRepo(app *firebase.App) ImageRepo {
	return &imageRepo{
		app: app,
	}
}

func (i *imageRepo) UploadImage(ctx context.Context, arg *domain.UploadImageParam) (string, string, error) {
	defer newrelic.FromContext(ctx).StartSegment("imageRepo-Upload").End()
	client, err := i.app.Storage(context.Background())
	if err != nil {
		return "", "", fmt.Errorf("storage.Client: %v", err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return "", "", fmt.Errorf("DefaultBucket: %v", err)
	}

	obj := bucket.Object(arg.Key)
	wc := obj.NewWriter(ctx)
	wc.ContentType = arg.ContentType

	if _, err := wc.Write(arg.Body); err != nil {
		return "", "", fmt.Errorf("createFile:write %v: %v", arg.Key, err)
	}
	if err := wc.Close(); err != nil {
		return "", "", fmt.Errorf("createFile:close %v: %v", arg.Key, err)
	}
	downloadURL, err := bucket.SignedURL(obj.ObjectName(), &storage.SignedURLOptions{
		Expires: time.Now().AddDate(100, 0, 0),
		Method:  "GET",
	})
	if err != nil {
		return "", "", fmt.Errorf("downloadURL :%v", err)
	}

	return downloadURL, arg.Key, nil
}

func (i *imageRepo) DeleteImage(ctx context.Context, key string) error {
	defer newrelic.FromContext(ctx).StartSegment("imageRepo-Delete").End()
	client, err := i.app.Storage(context.Background())
	if err != nil {
		return fmt.Errorf("storage.Client: %v", err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return fmt.Errorf("DefaultBucket: %v", err)
	}

	if err := bucket.Object(key).Delete(ctx); err != nil {
		return err
	}
	return nil
}
