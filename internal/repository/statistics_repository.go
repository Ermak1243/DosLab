package repository

import (
	"context"
	"log"
	"main/domain"
	"main/internal/database/postgres"
	"main/internal/services/mistakes"
	"time"
)

type statisticsRepository struct {
	database postgres.Database
}

func NewStatisticsRepository(db postgres.Database, timeout time.Duration) domain.StatisticsRepository {
	return &statisticsRepository{
		database: db,
	}
}

func (sr *statisticsRepository) CreateOrUpdate(ctx context.Context, statistics []domain.Statistics) error {
	_, err := sr.database.WithContext(ctx).Model(&statistics).OnConflict("(post_id, word) DO UPDATE").Insert()
	if err != nil {
		log.Println(mistakes.ErrDatabase(err, sr.CreateOrUpdate))
	}

	return err
}

func (sr *statisticsRepository) Fetch(ctx context.Context, postId string) ([]domain.Statistics, error) {
	var statistics []domain.Statistics

	err := sr.database.WithContext(ctx).Model(&statistics).Where("post_id = ?", postId).Order("count DESC").Select()
	if err != nil {
		log.Println(mistakes.ErrDatabase(err, sr.Fetch))
	}

	return statistics, err
}
