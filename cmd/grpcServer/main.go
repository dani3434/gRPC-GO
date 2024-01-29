package main

import (
	"database/sql"
	 _ "github.com/mattn/go-sqlite3"
	"github.com/devfullcycle/goexpert/tree/main/14-gRPC/internal/database"
	"github.com/devfullcycle/goexpert/tree/main/14-gRPC/internal/service"
	 pb "github.com/devfullcycle/goexpert/tree/main/14-gRPC/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	db,err := sql.Open("sqlite3","./db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
  reflection.Register(grpcServer)

	list,err := net.Listen("tcp",":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(list); err != nil {
		panic(err)
	}
}