package main

import (
	"github/galvizlaura69/backGo/controller"
	"github/galvizlaura69/backGo/model"
	"log"
	"net/http"
)

func main() {
	if err := model.InitDB(); err != nil {
		log.Fatal("Error al conectar a MongoDB:", err)
	}

	http.HandleFunc("/usuarios", controller.HandleUsuarios)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
