package domain

import "context"

type Statistics struct {
	PostId int    `json:"post_id" pg:",unique:id_word"`
	Word   string `json:"word" pg:",unique:id_word"`
	Count  int    `json:"count"`
}

type StatisticsRepository interface {
	CreateOrUpdate(ctx context.Context, statistics []Statistics) error
	Fetch(ctx context.Context, postId string) ([]Statistics, error)
}

type StatisticsUsecase interface {
	CreateOrUpdate(ctx context.Context, statistics []Statistics) error
	Fetch(ctx context.Context, postId string) ([]Statistics, error)
}
