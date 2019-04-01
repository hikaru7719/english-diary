package main

import (
	"context"
	"fmt"
	pb "github.com/hikaru7719/english-diary/proto"
	"google.golang.org/grpc"
	"time"
)

func main() {
	var opt []grpc.DialOption
	opt = append(opt, grpc.WithInsecure())
	conn, err := grpc.Dial(":8080", opt...)
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	client := pb.NewEnglishDiaryClient(conn)
	PrintGetDiary(client)
}

func PrintGetDiary(client pb.EnglishDiaryClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request := &pb.GetDiaryRequest{
		Id: 1,
	}
	diary, err := client.GetDiary(ctx, request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(diary)
}
