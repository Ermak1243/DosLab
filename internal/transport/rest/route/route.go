package route

import (
	"main/domain"
	"main/internal/config"
	"main/internal/database/postgres"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *config.Env, statisticsUsecase domain.StatisticsUsecase, db postgres.Database, gin *gin.Engine, timeout time.Duration) {
	publicRouter := gin.Group("")
	NewStatisticsRouter(timeout, statisticsUsecase, db, publicRouter)
}
