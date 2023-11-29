package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	member "github.com/Hack-Hack-geek-Vol10/services/member-service/api/v1"
	"github.com/Hack-Hack-geek-Vol10/services/member-service/cmd/config"
	"github.com/Hack-Hack-geek-Vol10/services/member-service/internal/infra"
	"github.com/Hack-Hack-geek-Vol10/services/member-service/internal/usecase"
	"github.com/murasame29/db-conn/sqldb/postgres"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn := postgres.NewConnection(
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.DBName,
		config.Config.Database.SSLMode,
		config.Config.Database.ConnectAttempts,
		config.Config.Database.ConnectWaitTime,
		config.Config.Database.ConnectBlocks,
	)
	defer conn.Close(ctx)

	db, err := conn.Connection()
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	member.RegisterMemberServiceServer(s, usecase.NewMemberService(infra.NewMemberRepo(db)))

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
