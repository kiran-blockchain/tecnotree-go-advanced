package database

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseHandler struct {
    client *mongo.Client
    db     *mongo.Database
}

func NewDatabaseHandler() (*DatabaseHandler, error) {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return nil, err
    }

    return &DatabaseHandler{
        client: client,
        db:     client.Database("mydb"),
    }, nil
}

func (dh *DatabaseHandler) FindDocument(collectionName string, filter interface{}) (*mongo.SingleResult, error) {
    collection := dh.db.Collection(collectionName)
    return collection.FindOne(context.TODO(), filter), nil
}
