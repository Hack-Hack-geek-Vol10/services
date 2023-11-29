package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Connect(path string) (*firebase.App, error) {
	config := &firebase.Config{
		StorageBucket: "geek-vol10.appspot.com",
	}

	serviceAccount := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(context.Background(), config, serviceAccount)
	if err != nil {
		return nil, err
	}

	return app, nil
}
