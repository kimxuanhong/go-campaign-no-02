package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

//go:generate mockgen -source=mongo_config.go -destination=mocks/mongo_config_mock.go -package=mocks

type MongoDB interface {
	GetConnection() *mongo.Client
	GetCollection(collectionName string) *mongo.Collection
}

type MongoDBImpl struct {
	db *mongo.Client
}

var instanceMongoDB *MongoDBImpl

func MongoDBInstance() *MongoDBImpl {
	if instanceMongoDB == nil {
		instanceMongoDB = &MongoDBImpl{
			db: connectMongoDB(),
		}
	}
	return instanceMongoDB
}

func (r *MongoDBImpl) GetConnection() *mongo.Client {
	if r.db == nil {
		r.db = connectMongoDB()
	}
	return r.db
}

func connectMongoDB() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

func (r *MongoDBImpl) GetCollection(collectionName string) *mongo.Collection {
	var collection = r.db.Database(os.Getenv("MONGO_DBNAME")).Collection(collectionName)
	return collection
}
