package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github/galvizlaura69/backGo/model"

	"go.mongodb.org/mongo-driver/bson"
)

type ControllerInsertar struct {
	Model *model.Model
}

func NewControllerInsertar(model *model.Model) *ControllerInsertar {
	return &ControllerInsertar{
		Model: model,
	}
}

// User represents the structure of the user document
type UserInsert struct {
	Documento       string `json:"documento"`
	Email           string `json:"email"`
	PrimerNombre    string `json:"primerNombre"`
	SegundoNombre   string `json:"segundoNombre"`
	PrimerApellido  string `json:"primerApellido"`
	SegundoApellido string `json:"segundoApellido"`
	Telefono        string `json:"telefono"`
}

func (c *ControllerInsertar) InsertUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user UserInsert
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	document := bson.D{
		{Key: "documento", Value: user.Documento},
		{Key: "primerNombre", Value: user.PrimerNombre},
		{Key: "segundoNombre", Value: user.SegundoNombre},
		{Key: "primerApellido", Value: user.PrimerApellido},
		{Key: "segundoApellido", Value: user.SegundoApellido},
		{Key: "email", Value: user.Email},
		{Key: "telefono", Value: user.Telefono},
	}

	_, err := c.Model.Collection.InsertOne(context.TODO(), document)
	if err != nil {
		log.Println("Error inserting document into the database:", err)
		http.Error(w, "Error inserting document", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println("Error encoding response to JSON:", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
