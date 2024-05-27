package route

import (
	"main/domain"
	"main/internal/database/postgres"
	"main/internal/transport/rest/controller"
	"time"

	"github.com/gin-gonic/gin"
)

func NewStatisticsRouter(timeout time.Duration, statisticsUsecase domain.StatisticsUsecase, db postgres.Database, group *gin.RouterGroup) {
	sc := controller.StatisticsController{
		StatisticsUsecase: statisticsUsecase,
	}

	group.GET("/post/:postId/comments/statistics", sc.Statistics)
}
