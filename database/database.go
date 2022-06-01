package database

import (
	"context"
	"echo-app/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var db *mongo.Database

func Connect() {
	var env = config.GetEnv()

	// Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.Database.URI))
	if err != nil {
		fmt.Println(err.Error())
	}

	db = client.Database(env.Database.Name)
	fmt.Println("Database connected to", env.Database.Name)
}
