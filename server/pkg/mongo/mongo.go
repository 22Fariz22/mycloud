package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() *mongo.Database {
	fmt.Println("in initDB")
	client, err := mongo.Connect(context.Background(),options.Client().ApplyURI(viper.GetString("mongo.uri")))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err, " in error mongo client connect")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("log fatal in ping mongo")
	}

	return client.Database(viper.GetString("mongo.name"))
}
