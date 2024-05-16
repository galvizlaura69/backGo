package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://galvizlaura69:Canela30+@cluster0.cegzqm2.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error al conectar a MongoDB:", err)
	}
	defer client.Disconnect(context.TODO())

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Error al hacer ping a MongoDB:", err)
	}
	fmt.Println("Conexión exitosa a MongoDB")

	collection := client.Database("test").Collection("usuarios")

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		cursor, err := collection.Find(context.TODO(), bson.D{})
		if err != nil {
			log.Println("Error al buscar usuarios en la base de datos:", err)
			http.Error(w, "Error fetching data", http.StatusInternalServerError)
			return
		}
		defer cursor.Close(context.TODO())

		var results []bson.M
		if err := cursor.All(context.TODO(), &results); err != nil {
			log.Println("Error al decodificar datos de la base de datos:", err)
			http.Error(w, "Error decoding data", http.StatusInternalServerError)
			return
		}

		log.Println("Número de resultados:", len(results))

		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(results); err != nil {
			log.Println("Error al codificar datos a JSON:", err)
			http.Error(w, "Error encoding data to JSON", http.StatusInternalServerError)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
