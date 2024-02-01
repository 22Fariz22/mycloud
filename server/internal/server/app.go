package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/22Fariz22/mycloud/server/internal/auth"
	repoMongo "github.com/22Fariz22/mycloud/server/internal/auth/repository/mongo"
	authusecase "github.com/22Fariz22/mycloud/server/internal/auth/usecase"
	"github.com/22Fariz22/mycloud/server/internal/file"
	repoPG "github.com/22Fariz22/mycloud/server/internal/file/repository/postgres"
	fileusecase "github.com/22Fariz22/mycloud/server/internal/file/usecase"
	"github.com/22Fariz22/mycloud/server/pkg/logger"
	"github.com/22Fariz22/mycloud/server/pkg/postgres"

	"github.com/spf13/viper"
)

type App struct {
	httpServer *http.Server

	fileUC file.UseCase
	authUC auth.UseCase

	logger logger.Logger
}

func NewApp() *App {
	db := InitDB()

	appLogger := logger.NewApiLogger()

	appLogger.InitLogger()
	// appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL)

	psqlDB, err := postgres.NewPsqlDB()
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
	} else {
		appLogger.Infof("Postgres connected, Status: %#v", psqlDB.Stats())
	}
	defer psqlDB.Close()

	userRepo := repoMongo.NewUserRepository(db, viper.GetString("mongo.user_collection"))
	fileRepo := repoPG.NewFileRepository(psqlDB)

	return &App{
		fileUC: fileusecase.NewFileUC(fileRepo),
		authUC: authusecase.NewAuthUseCase(
			userRepo,
			viper.GetString("auth.hash_salt"),
			[]byte(viper.GetString("auth.signing_key")),
			viper.GetDuration("auth.token_ttl"),
		),
	}

}

func (a *App) Run(port string) error {

	return nil
}

func InitDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.uri")))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err, "error mongo client connect")
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("log fatal")
	}

	return client.Database(viper.GetString("mongo.name"))
}
