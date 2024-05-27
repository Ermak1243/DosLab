package app

import (
	"context"
	"fmt"
	"main/domain"
	"main/internal/config"
	"main/internal/database/postgres"
	"main/internal/models"
	"main/internal/repository"
	"main/internal/services"
	"main/internal/transport/rest/route"
	"main/internal/usecase"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map/v2"
)

var ctx = context.Background()

type Application struct {
	Env               *config.Env
	Postgres          postgres.Database
	StatisticsUsecase domain.StatisticsUsecase
}

func Run() {
	app := &Application{}
	app.Env = config.NewEnv("../../configs/.env")
	app.Postgres = postgres.NewPostgresDatabase(app.Env)
	defer app.CloseDBConnection()

	timeout := time.Duration(app.Env.ContextTimeout) * time.Second

	sr := repository.NewStatisticsRepository(app.Postgres, timeout)
	app.StatisticsUsecase = usecase.NewStatisticsUsecase(sr, timeout)

	placeholder := services.NewPlaceholder()

	go func() {
		for {
			comments := placeholder.Comments()
			numberOfMeetings := cmap.New[cmap.ConcurrentMap[string, int]]()
			var dataToInsert []domain.Statistics

			var wg sync.WaitGroup

			for _, comment := range comments {
				wg.Add(1)

				go func(comment models.Comment) {
					defer wg.Done()

					words := strings.Split(strings.Replace(comment.Body, "\n", " ", -1), " ")

					numberOfMeetings.SetIfAbsent(strconv.Itoa(comment.PostID), cmap.New[int]())

					for _, word := range words {
						wordsInOnePost, _ := numberOfMeetings.Get(strconv.Itoa(comment.PostID))

						wordCounter, _ := wordsInOnePost.Get(word)
						wordsInOnePost.Set(word, wordCounter+1)
					}
				}(comment)
			}

			wg.Wait()

			for data := range numberOfMeetings.IterBuffered() {
				/*Здесь оптимальным будет завернуть блок кода на каждой итерации в горутину, как в предыдущем цикле с использованием sync.WaitGroup,
				но, видимо, из-за внутренней реализации пакета "cmap", код работает некорректно. Логику выполнения можно было реализовать и по-другому - без hashMap.
				*/
				postId, err := strconv.Atoi(data.Key)
				if err != nil {
					fmt.Println(err)
				}

				for word := range data.Val.IterBuffered() {
					statistic := domain.Statistics{
						PostId: postId,
						Word:   word.Key,
						Count:  word.Val,
					}

					dataToInsert = append(dataToInsert, statistic)
				}
			}

			app.StatisticsUsecase.CreateOrUpdate(ctx, dataToInsert)

			time.Sleep(5 * time.Minute)
		}
	}()

	gin := gin.Default()

	route.Setup(app.Env, app.StatisticsUsecase, app.Postgres, gin, timeout)

	gin.Run(app.Env.ServerAddress)
}

func (app *Application) CloseDBConnection() {
	postgres.ClosePostgresDBConnection(app.Postgres)
}
