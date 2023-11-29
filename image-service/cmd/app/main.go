package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	image "github.com/Hack-Hack-geek-Vol10/services/image-service/api/v1"
	"github.com/Hack-Hack-geek-Vol10/services/image-service/cmd/config"
	"github.com/Hack-Hack-geek-Vol10/services/image-service/internal/infra"
	"github.com/Hack-Hack-geek-Vol10/services/image-service/internal/usecase"
	"github.com/Hack-Hack-geek-Vol10/services/pkg/driver/firebase"
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

	app, err := firebase.Connect("./sa.json")
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
