package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	firebase "firebase.google.com/go"
	image "github.com/schema-creator/services/image-service/api/v1"
	"github.com/schema-creator/services/image-service/cmd/config"
	"github.com/schema-creator/services/image-service/internal/infra"
	"github.com/schema-creator/services/image-service/internal/usecase"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	config.LoadEnv()
}

func main() {

	listener, err := net.Listen("tcp", config.Config.Server.ServerAddr)
	if err != nil {
		panic(err)
	}

	serviceAccount := option.WithCredentialsFile("./sa.json")
	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		StorageBucket: config.Config.Firebase.Bucket,
	}, serviceAccount)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	image.RegisterImageServiceServer(s, usecase.NewImageService(infra.NewImageRepo(app)))

	reflection.Register(s)

	// 3. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", config.Config.Server.ServerAddr)
		s.Serve(listener)
	}()

	// 4.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
