package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Hack-Hack-geek-Vol10/services/cmd/config"
	image "github.com/Hack-Hack-geek-Vol10/services/pkg/grpc/image-service/v1"
	"github.com/Hack-Hack-geek-Vol10/services/src/driver/firebase"
	"github.com/Hack-Hack-geek-Vol10/services/src/services"
	"github.com/Hack-Hack-geek-Vol10/services/src/storages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.LoadEnv()

	listener, err := net.Listen("tcp", config.Config.Server.Imageaddr)
	if err != nil {
		panic(err)
	}

	app, err := firebase.Connect("./sa.json")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	image.RegisterImageServiceServer(s, services.NewImageService(storages.NewImageRepo(app)))

	reflection.Register(s)

	// 3. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", config.Config.Server.Imageaddr)
		s.Serve(listener)
	}()

	// 4.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
