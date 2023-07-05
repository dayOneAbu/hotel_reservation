package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LoadEnv(name string) string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	envVarName := os.Getenv(name)

	return envVarName
}
func ConnectTOMongo() *mongo.Client {
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("No .env file found")
	// }
	uri := LoadEnv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	return client
}
