package main

import (
	"google.golang.org/grpc"
	"net"

	"github.com/hikaru7719/english-diary/interface/handler"
	pb "github.com/hikaru7719/english-diary/proto"
)

func Serve() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterEnglishDiaryServer(grpcServer, handler.NewEnglishDiaryService())
	grpcServer.Serve(lis)
}
