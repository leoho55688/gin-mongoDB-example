package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"backend/mongo"
)

func NewMongoDatabase(env *Env) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbUser := env.DBUser
	dbPass := env.DBPass

	mongodbURI := fmt.Sprintf("%s+srv://%s:%s@fyp.2tvkfky.mongodb.net/?retryWrites=true&w=majority", dbHost, dbUser, dbPass)

	client, err := mongo.NewClient(ctx, mongodbURI)
	if err != nil {
		log.Fatal(err)
	} 

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println(("Connection to MongoDB closed."))
}