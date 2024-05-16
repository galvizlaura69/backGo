package model

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	Collection *mongo.Collection
}

func NewModel() *Model {
	return &Model{}
}

func (m *Model) InitDB() error {
	clientOptions := options.Client().ApplyURI("mongodb+srv://galvizlaura69:Canela30+@cluster0.cegzqm2.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	m.Collection = client.Database("test").Collection("usuarios")
	return nil
}
