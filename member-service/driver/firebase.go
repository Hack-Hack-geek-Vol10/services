package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/Hack-Hack-geek-Vol10/services/image-service/cmd/config"
	"google.golang.org/api/option"
)

func Connect(path string) (*firebase.App, error) {
	config := &firebase.Config{
		StorageBucket: config.Config.Firebase.Bucket,
	}

	serviceAccount := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(context.Background(), config, serviceAccount)
	if err != nil {
		return nil, err
	}

	return app, nil
}
