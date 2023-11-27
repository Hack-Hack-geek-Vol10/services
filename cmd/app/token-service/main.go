package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Hack-Hack-geek-Vol10/services/cmd/config"
	token "github.com/Hack-Hack-geek-Vol10/services/pkg/grpc/token-service/v1"
	"github.com/Hack-Hack-geek-Vol10/services/pkg/maker"
	"github.com/Hack-Hack-geek-Vol10/services/src/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.LoadEnv()

	listener, err := net.Listen("tcp", config.Config.Server.TokenAddr)
	if err != nil {
		panic(err)
	}

	tokenMaker, err := maker.NewPasetoMaker(config.Config.Token.SymmetricKey)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	token.RegisterTokenServiceServer(s, services.NewTokenService(tokenMaker))

	reflection.Register(s)

	// 3. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", config.Config.Server.TokenAddr)
		s.Serve(listener)
	}()

	// 4.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
