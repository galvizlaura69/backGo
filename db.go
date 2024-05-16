package main

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
    clientOptions := options.Client().ApplyURI("mongodb+srv://galvizlaura69:Canela30+@cluster0.cegzqm2.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
    var err error
    client, err = mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Conexión establecida con MongoDB")
}

func getCollection(collectionName string) *mongo.Collection {
    return client.Database("BACKENDBD").Collection(collectionName)
}
