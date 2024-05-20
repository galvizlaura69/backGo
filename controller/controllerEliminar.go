package controller

import (
	"context"
	"log"
	"net/http"

	"github/galvizlaura69/backGo/model"

	"go.mongodb.org/mongo-driver/bson"
)

type ControllerEliminar struct {
	Model *model.Model
}

func NewControllerEliminar(model *model.Model) *ControllerEliminar {
	return &ControllerEliminar{
		Model: model,
	}
}

func (c *ControllerEliminar) EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	documento := r.URL.Query().Get("documento")
	if documento == "" {
		http.Error(w, "Missing documento parameter", http.StatusBadRequest)
		return
	}

	filter := bson.D{{"documento", documento}}

	result, err := c.Model.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("Error deleting document:", err)
		http.Error(w, "Error deleting document", http.StatusInternalServerError)
		return
	}

	log.Println("Número de documentos eliminados:", result.DeletedCount)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Document deleted successfully"))
}
