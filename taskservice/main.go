package main

import (
	"database/sql"
	"log"
	"net"
	"os"
	"time"

	"taskservice/api"
	"taskservice/proto"
	"taskservice/repository"
	"taskservice/service"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load .env file: %v", err)
	}

	cfg := mysql.Config{
		User:   os.Getenv("TEST_DB_USER"),
		Passwd: os.Getenv("TEST_DB_PASSWD"),
		Net:    os.Getenv("TEST_DB_NET"),
		Addr:   os.Getenv("TEST_DB_ADDR"),
		DBName: os.Getenv("TEST_DB_NAME"),
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(120 * time.Second)

	taskRepo := repository.NewTaskRepo(db)
	taskService := service.NewTaskService(taskRepo)
	taskController := api.NewTaskController(taskService)

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen on port 8000: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterTaskServiceGRPCServer(grpcServer, taskController)

	log.Println("gRPC server running on port 8000")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
