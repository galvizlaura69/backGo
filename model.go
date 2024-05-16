package model

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func InitDB() error {
	clientOptions := options.Client().ApplyURI("mongodb+srv://galvizlaura69:Canela30+@cluster0.cegzqm2.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	Collection = client.Database("test").Collection("usuarios")
	return nil
}
