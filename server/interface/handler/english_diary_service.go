package handler

import (
	"context"
	pb "github.com/hikaru7719/english-diary/proto"
)

type englishDiaryService struct {
}

func NewEnglishDiaryService() *englishDiaryService {
	return &englishDiaryService{}
}

func (e *englishDiaryService) GetDiary(context.Context, *pb.GetDiaryRequest) (*pb.Diary, error) {
	diary := &pb.Diary{
		Id:    1,
		Title: "sample",
		Body:  "sample diary",
	}

	return diary, nil
}
