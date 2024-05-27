package controller

import (
	"main/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatisticsController struct {
	StatisticsUsecase domain.StatisticsUsecase
}

func (sc *StatisticsController) Statistics(c *gin.Context) {
	statistics, err := sc.StatisticsUsecase.Fetch(c, c.Param("postId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, statistics)
	}

	c.JSON(http.StatusOK, statistics)
}
