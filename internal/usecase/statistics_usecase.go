package usecase

import (
	"context"
	"log"
	"main/domain"

	"time"
)

type statisticsUsecase struct {
	statisticsRepository domain.StatisticsRepository
	contextTimeout       time.Duration
}

func NewStatisticsUsecase(statisticsRepository domain.StatisticsRepository, timeout time.Duration) domain.StatisticsUsecase {
	return &statisticsUsecase{
		statisticsRepository: statisticsRepository,
		contextTimeout:       timeout,
	}
}

func (su *statisticsUsecase) CreateOrUpdate(ctx context.Context, statistics []domain.Statistics) error {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()

	err := su.statisticsRepository.CreateOrUpdate(ctx, statistics)
	if err != nil {
		log.Println(err)

		return err
	}

	return err
}

func (su *statisticsUsecase) Fetch(ctx context.Context, postId string) ([]domain.Statistics, error) {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()

	statistics, err := su.statisticsRepository.Fetch(ctx, postId)
	if err != nil {
		log.Println(err)

		return []domain.Statistics{}, err
	}

	return statistics, err
}
