package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/22Fariz22/mycloud/server/config"
	"github.com/22Fariz22/mycloud/server/pkg/logger"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

// NewPsqlDB Return new Postgresql db instance
func NewPsqlDB(c *config.Config, logger logger.Logger,) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.Postgres.PostgresqlHost,
		c.Postgres.PostgresqlPort,
		c.Postgres.PostgresqlUser,
		c.Postgres.PostgresqlDbname,
		c.Postgres.PostgresqlPassword,
	)
	log.Println("pkg.postgres.NewPsqlDB()--dataSourceName: ",dataSourceName)

	db, err := sqlx.Connect(c.Postgres.PgDriver, dataSourceName)
	if err != nil {
		logger.Error("error in pkg.postrgres.NewPsqlDB()--sqlx.Connect()")
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	if err = db.Ping(); err != nil {
		logger.Error("error in pkg.postrgres.NewPsqlDB()--Ping()")
		return nil, err
	}

	return db, nil
}
