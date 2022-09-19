package configs

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	DBClient   *mongo.Client
	DBInstance *mongo.Database
)

func DBContext() context.Context {
	return context.Background()
}

func InitMongoDB() {
	client, err := mongo.Connect(DBContext(),
		options.Client().ApplyURI(Env.MONGO_URI),
		options.Client().SetTimeout(time.Second*10),
	)
	if err != nil {
		panic(fmt.Sprintf("Connect database failed: %s", err.Error()))
	}

	err = client.Ping(DBContext(), readpref.Primary())
	if err != nil {
		panic(fmt.Sprintf("Connect database failed: %s", err.Error()))
	}

	fmt.Printf("Database URI: %s \nDatabase: %s \n", Env.MONGO_URI, Env.MONGO_DATABASE_NAME)

	DBClient = client
	DBInstance = client.Database(Env.MONGO_DATABASE_NAME)
}
