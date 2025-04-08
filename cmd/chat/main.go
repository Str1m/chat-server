package main

import (
	"chat-server/internal/config"
	"chat-server/internal/config/env"
	"chat-server/internal/grpc/chat"
	"chat-server/internal/storage/postgres"
	desc "chat-server/pkg/chat_v1"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")

	flag.Parse()

	config.MustLoad(configPath)

	storageConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to load storage config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to load grpc config: %v", err)
	}

	fmt.Println(storageConfig.DSN())

	_, err = postgres.New(storageConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	l, err := net.Listen("tcp", grpcConfig.Addr())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	desc.RegisterChatServer(s, &chat.Server{})

	log.Printf("server listening: %s\n", grpcConfig.Addr())

	if err = s.Serve(l); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
