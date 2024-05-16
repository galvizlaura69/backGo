package main

import (
	"log"
	"net/http"

	"github.com/galvizlaura69/backGo/model"
)

func main() {
	if err := model.InitDB(); err != nil {
		log.Fatal("Error al conectar a MongoDB:", err)
	}

	http.HandleFunc("/usuarios", controller.HandleUsuarios)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
