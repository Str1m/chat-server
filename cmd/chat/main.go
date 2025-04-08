package main

import (
	"chat-server/internal/config"
	"chat-server/internal/config/env"
	"chat-server/internal/models"
	"chat-server/internal/storage/postgres"
	"context"
	"flag"
	"fmt"
	"log"
	"time"
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

	// grpcConfig, err := env.NewGRPCConfig()
	// if err != nil {
	// 	log.Fatalf("failed to load grpc config: %v", err)
	// }

	fmt.Println(storageConfig.DSN())

	ctx := context.Background()

	store, err := postgres.New(storageConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	usernames := []string{"alice@example.com", "bob@example.com"}
	chatID, err := store.CreateChat(ctx, usernames)
	if err != nil {
		log.Fatalf("cannot create chat: %v", err)
	}
	fmt.Println("Chat created with ID:", chatID)

	msg := models.Message{
		ChatID:    chatID,
		Sender:    "alice@example.com",
		Text:      "Hello, Bob!",
		Timestamp: time.Now(),
	}

	if err := store.SendMessage(msg); err != nil {
		log.Fatalf("cannot send message: %v", err)
	}
	fmt.Println("Message sent!")

	messages, err := store.GetMessages(chatID)
	if err != nil {
		log.Fatalf("cannot get messages: %v", err)
	}

	fmt.Println("Messages in chat:")
	for _, m := range messages {
		fmt.Printf("[%s] %s: %s\n", m.Timestamp.Format(time.RFC822), m.Sender, m.Text)
	}

	// l, err := net.Listen("tcp", grpcConfig.Addr())
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// s := grpc.NewServer()
	// reflection.Register(s)

	// desc.RegisterChatServer(s, &chat.Server{})

	// log.Printf("server listening: %s\n", grpcConfig.Addr())

	// if err = s.Serve(l); err != nil {
	// 	log.Fatalf("failed to server: %v", err)
	// }
}
