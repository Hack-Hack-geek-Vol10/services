package storages

import (
	"bytes"
	"context"

	"github.com/Hack-Hack-geek-Vol10/services/cmd/config"
	"github.com/Hack-Hack-geek-Vol10/services/src/domain"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type imageRepo struct {
	c *s3.Client
}

type ImageRepo interface {
	UploadImage(ctx context.Context, arg *domain.UploadImageParam) (string, string, error)
	DeleteImage(ctx context.Context, key string) error
}

func NewImageRepo(c *s3.Client) ImageRepo {
	return &imageRepo{
		c: c,
	}
}

func (i *imageRepo) UploadImage(ctx context.Context, arg *domain.UploadImageParam) (string, string, error) {
	result, err := manager.NewUploader(i.c).Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(config.Config.S3.CfBucket),
		Key:         aws.String(arg.Key),
		Body:        bytes.NewReader(arg.Body),
		ContentType: aws.String(arg.ContentType),
	})
	if err != nil {
		return "", "", err
	}

	return result.Location, *result.Key, nil
}

func (i *imageRepo) DeleteImage(ctx context.Context, key string) error {
	_, err := i.c.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(config.Config.S3.CfBucket),
		Key:    aws.String(key),
	})
	return err
}
