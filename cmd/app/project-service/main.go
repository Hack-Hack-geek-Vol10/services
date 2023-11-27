package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/Hack-Hack-geek-Vol10/services/cmd/config"
	project "github.com/Hack-Hack-geek-Vol10/services/pkg/grpc/project-service/v1"
	"github.com/Hack-Hack-geek-Vol10/services/src/driver/postgres"
	"github.com/Hack-Hack-geek-Vol10/services/src/services"
	"github.com/Hack-Hack-geek-Vol10/services/src/storages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.LoadEnv()

	listener, err := net.Listen("tcp", config.Config.Server.ProjectAddr)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn := postgres.NewConnection()
	defer conn.Close(ctx)

	db, err := conn.Connection()
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	project.RegisterProjectServiceServer(s, services.NewProjectService(storages.NewProjectRepo(db)))

	reflection.Register(s)

	// 3. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", config.Config.Server.ProjectAddr)
		s.Serve(listener)
	}()

	// 4.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
