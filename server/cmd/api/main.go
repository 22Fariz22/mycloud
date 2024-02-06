package main

import (
	"fmt"
	"log"

	"github.com/22Fariz22/mycloud/server/config"
	"github.com/22Fariz22/mycloud/server/internal/server"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("main")
	if err := config.Init(); err != nil {
		log.Fatalf("%s in main-config.init", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s in app.run", err.Error())
	}
}
