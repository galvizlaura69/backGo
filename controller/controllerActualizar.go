package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github/galvizlaura69/backGo/model"

	"go.mongodb.org/mongo-driver/bson"
)

type ControllerActualizar struct {
	Model *model.Model
}

func NewControllerActualizar(model *model.Model) *ControllerActualizar {
	return &ControllerActualizar{
		Model: model,
	}
}

// UpdateDocument represents the structure of the update document
type UpdateDocument struct {
	Documento       string `json:"documento"`
	Email           string `json:"email"`
	PrimerNombre    string `json:"primerNombre"`
	SegundoNombre   string `json:"segundoNombre"`
	PrimerApellido  string `json:"primerApellido"`
	SegundoApellido string `json:"segundoApellido"`
	Telefono        string `json:"telefono"`
}

func (c *ControllerActualizar) ActualizarUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	documento := r.URL.Query().Get("documento")
	if documento == "" {
		http.Error(w, "Missing documento parameter", http.StatusBadRequest)
		return
	}

	var update UpdateDocument
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	filter := bson.D{{"documento", documento}}
	updateDoc := bson.D{
		{"$set", bson.D{
			{"primerNombre", update.PrimerNombre},
			{"segundoNombre", update.SegundoNombre},
			{"primerApellido", update.PrimerApellido},
			{"segundoApellido", update.SegundoApellido},
			{"email", update.Email},
			{"telefono", update.Telefono},
		}},
	}

	result, err := c.Model.Collection.UpdateOne(context.TODO(), filter, updateDoc)
	if err != nil {
		log.Println("Error updating document:", err)
		http.Error(w, "Error updating document", http.StatusInternalServerError)
		return
	}

	log.Println("Número de documentos actualizados:", result.ModifiedCount)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Document updated successfully"))
}
