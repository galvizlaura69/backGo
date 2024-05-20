package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github/galvizlaura69/backGo/model"

	"go.mongodb.org/mongo-driver/bson"
)

type ControllerConsultar struct {
	Model *model.Model
}

func NewControllerConsultar(model *model.Model) *ControllerConsultar {
	return &ControllerConsultar{
		Model: model,
	}
}

func (c *ControllerConsultar) HandleUsuarios(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	documento := r.URL.Query().Get("documento")
	if documento == "" {
		http.Error(w, "Missing documento parameter", http.StatusBadRequest)
		return
	}

	filter := bson.D{{"documento", documento}}
	cursor, err := c.Model.Collection.Find(context.TODO(), filter)
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
