package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	config "github.com/klan300/exceed17/config"
)

func DatabaseConnect() ( context.Context, *mongo.Database)  {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GoDotEnvVariable("DATABASEURI")))

	if err != nil {
        log.Fatal(err)
	}
	
	ctx, _ := context.WithTimeout(context.Background(),1*time.Minute)
	
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(config.GoDotEnvVariable("DATABASE"))

	return ctx, database
}