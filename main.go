package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	"github/galvizlaura69/backGo/controller"
	"github/galvizlaura69/backGo/model"
)

func main() {
	app := NewApp()
	if err := app.Init(); err != nil {
		log.Fatalf("Error initializing app: %s", err)
	}

	corsMiddleware := cors.Default().Handler

	http.Handle("/consultar", corsMiddleware(http.HandlerFunc(app.ControllerConsultar.HandleUsuarios)))
	http.Handle("/insertar", corsMiddleware(http.HandlerFunc(app.ControllerInsertar.InsertUser)))
	http.Handle("/actualizar", corsMiddleware(http.HandlerFunc(app.ControllerActualizar.ActualizarUsuario)))
	http.Handle("/eliminar", corsMiddleware(http.HandlerFunc(app.ControllerEliminar.EliminarUsuario)))

	log.Fatal(http.ListenAndServe(":8000", nil))
}

type App struct {
	Model                *model.Model
	ControllerConsultar  *controller.ControllerConsultar
	ControllerInsertar   *controller.ControllerInsertar
	ControllerActualizar *controller.ControllerActualizar
	ControllerEliminar   *controller.ControllerEliminar
}

func NewApp() *App {
	modelInstance := model.NewModel()
	return &App{
		Model:                modelInstance,
		ControllerConsultar:  controller.NewControllerConsultar(modelInstance),
		ControllerInsertar:   controller.NewControllerInsertar(modelInstance),
		ControllerActualizar: controller.NewControllerActualizar(modelInstance),
		ControllerEliminar:   controller.NewControllerEliminar(modelInstance),
	}
}

func (a *App) Init() error {
	if err := a.Model.InitDB(); err != nil {
		return err
	}
	return nil
}
