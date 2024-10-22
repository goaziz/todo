package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goaziz/todo/internal/todo" // Import your service package
	pb "github.com/goaziz/todo/proto/todo" // Adjust this import path based on your project structure
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

// @title ToDo Service API
// @version 1.0
// @description This is a ToDo service with gRPC and MongoDB
// @host localhost:8080
// @BasePath /v1
func main() {
	// Create a Gin router
	r := gin.Default()

	// Start the gRPC server
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		// Register the ToDoService with the gRPC server
		pb.RegisterToDoServiceServer(s, todo.NewToDoService()) // Use the imported function
		log.Printf("Starting gRPC server on port :%d", 50051)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start HTTP reverse proxy (gRPC gateway)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterToDoServiceHandlerFromEndpoint(context.Background(), mux, fmt.Sprintf("localhost:%d", 50051), opts)
	if err != nil {
		log.Fatalf("failed to start HTTP reverse proxy: %v", err)
	}

	// Mount the mux to the Gin router
	r.Any("/v1/todo/*any", gin.WrapH(mux))

	log.Printf("Starting HTTP server on port :%d", 8080)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
