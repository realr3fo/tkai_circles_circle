package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"github.com/realr3fo/tkai_circles_circle/app"
	"github.com/realr3fo/tkai_circles_circle/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/circle/area", controllers.GetCircleArea).Methods("POST")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
