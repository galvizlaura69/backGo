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

	http.Handle("/usuarios", corsMiddleware(http.HandlerFunc(app.Controller.HandleUsuarios)))

	log.Fatal(http.ListenAndServe(":8000", nil))
}

type App struct {
	Model      *model.Model
	Controller *controller.Controller
}

func NewApp() *App {
	modelInstance := model.NewModel()
	return &App{
		Model:      modelInstance,
		Controller: controller.NewController(modelInstance),
	}
}

func (a *App) Init() error {
	if err := a.Model.InitDB(); err != nil {
		return err
	}
	return nil
}
