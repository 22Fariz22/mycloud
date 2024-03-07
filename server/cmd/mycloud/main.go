package main

import (
	"log"
	"os"

	"github.com/22Fariz22/mycloud/server/config"
	"github.com/22Fariz22/mycloud/server/internal/app"
	"github.com/22Fariz22/mycloud/server/pkg/logger"
	"github.com/22Fariz22/mycloud/server/pkg/postgres"
	"github.com/22Fariz22/mycloud/server/pkg/redis"
	"github.com/22Fariz22/mycloud/server/pkg/utils"
)

func main() {
	log.Println("Starting service")

	configPath := utils.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config: %v", err)
	}

	appLogger := logger.NewAPILogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v",
		cfg.Server.AppVersion,
		cfg.Logger.Level,
		cfg.Server.Mode,
		cfg.Server.SSL,
	)
	appLogger.Infof("Success parsed config: %#v", cfg.Server.AppVersion)

	psqlDB, err := postgres.NewPsqlDB(cfg, appLogger)
	if err != nil {
		appLogger.Errorf("Postgresql init: %s", err)
	}
	defer psqlDB.Close()

	// if cfg.Postgres.EnableMigration{
	// 	m, err := migrate.New(
	// 	"file://migrations",
	// 	"postgres://postgres:postgres@localhost:5432/mycloud_db?sslmode=disable")
	// if err != nil {
	// 	appLogger.Errorf("err in migrate.New(): ",err)
	// }
	// if err := m.Up(); err != nil {
	// 	appLogger.Errorf("err in m.Up(): ",err)
	// }
	// }

	redisClient := redis.NewRedisClient(cfg)
	defer redisClient.Close()
	appLogger.Info("Redis connected")

	authServer := app.NewAuthServer(appLogger, cfg, psqlDB, redisClient)
	appLogger.Fatal(authServer.Run())
}
