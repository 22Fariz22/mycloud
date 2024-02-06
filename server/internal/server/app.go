package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/22Fariz22/mycloud/server/internal/auth"
	authhttp "github.com/22Fariz22/mycloud/server/internal/auth/delivery/http"
	repoMongo "github.com/22Fariz22/mycloud/server/internal/auth/repository/mongo"
	authusecase "github.com/22Fariz22/mycloud/server/internal/auth/usecase"
	filehttp "github.com/22Fariz22/mycloud/server/internal/file/delivery/http"

	"github.com/22Fariz22/mycloud/server/internal/file"
	repoPG "github.com/22Fariz22/mycloud/server/internal/file/repository/postgres"
	fileusecase "github.com/22Fariz22/mycloud/server/internal/file/usecase"
	"github.com/22Fariz22/mycloud/server/pkg/logger"
	mongodb "github.com/22Fariz22/mycloud/server/pkg/mongo"
	"github.com/22Fariz22/mycloud/server/pkg/postgres"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

type App struct {
	httpServer *http.Server

	fileUC file.UseCase
	authUC auth.UseCase

	//logger logger.Logger
}

func NewApp() *App {
	fmt.Println("in newApp")
	fmt.Println("viper get mongo uri :", viper.GetString("mongo.uri"))
	fmt.Println("viper get  :", viper.GetString("postgres.PostgresqlPort"))


	appLogger := logger.NewApiLogger()

	appLogger.InitLogger()
	// appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL)

  db := mongodb.InitDB()

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
	fmt.Println("Run")
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	authhttp.RegisterHTTPEndpoints(router, a.authUC)

	authMiddleware := authhttp.NewAuthMiddleware(a.authUC)
	api := router.Group("/api", authMiddleware)

	filehttp.RegisterHTTPEndpoints(api, a.fileUC)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
