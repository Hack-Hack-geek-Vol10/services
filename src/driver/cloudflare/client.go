package cloudflare

import (
	"context"
	"fmt"

	env "github.com/Hack-Hack-geek-Vol10/services/cmd/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type awsConfig struct {
	accountID        string
	accountKeyID     string
	accountKeySecret string
	s3Endpoint       string
}

func NewAwsConfig() *awsConfig {
	return &awsConfig{
		accountID:        env.Config.S3.CfAccountID,
		accountKeyID:     env.Config.S3.CfAccountKeyID,
		accountKeySecret: env.Config.S3.CfAccountKeySecret,
		s3Endpoint:       env.Config.S3.CfS3Endpoint,
	}
}

func (a *awsConfig) Client() (*s3.Client, error) {
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", a.accountID),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(a.accountKeyID, a.accountKeySecret, "")),
	)
	return s3.NewFromConfig(cfg), err
}
