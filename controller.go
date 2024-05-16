package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/galvizlaura69/backGo/model"
	"go.mongodb.org/mongo-driver/bson"
)

func HandleUsuarios(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cursor, err := model.Collection.Find(context.TODO(), bson.D{})
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
}
