package postgres

import (
	"fmt"
	"log"
	"time"

	"main/domain"
	"main/internal/config"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"golang.org/x/net/context"
)

type Database interface {
	String() string
	Close() error
	Options() *pg.Options
	WithTimeout(d time.Duration) *pg.DB
	WithParam(param string, value interface{}) *pg.DB
	Listen(ctx context.Context, channels ...string) *pg.Listener
	WithContext(ctx context.Context) *pg.DB
	Context() context.Context
	Model(model ...interface{}) *orm.Query
	Exec(query interface{}, params ...interface{}) (res pg.Result, err error)
}

func NewPostgresDatabase(env *config.Env) Database {
	var (
		dbUsername = env.PostgresDB.UserName
		dbPassword = env.PostgresDB.Password
		dbName     = env.PostgresDB.DbName
		dbHost     = env.PostgresDB.Host
		dbPort     = env.PostgresDB.Port
		dbAdress   = fmt.Sprintf("%s:%s", dbHost, dbPort)
	)

	db := pg.Connect(&pg.Options{
		User:     dbUsername,
		Password: dbPassword,
		Addr:     dbAdress,
		Database: dbName,
	})

	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}

	migrate(db)

	return db
}

func migrate(db Database) {
	db.Model(&domain.Statistics{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
}

func ClosePostgresDBConnection(con Database) {
	if con == nil {
		return
	}

	err := con.Close()
	if err != nil {
		log.Println(err)
	}

	log.Println("Connection to PostgresDB closed.")
}
