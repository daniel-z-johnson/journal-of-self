package main

import (
	"fmt"
	"net/http"

	"github.com/daniel-z-johnson/journal-of-self/config"
	"github.com/daniel-z-johnson/journal-of-self/controllers"
	"github.com/gorilla/mux"

	"github.com/daniel-z-johnson/journal-of-self/models"
)

func main() {
	fmt.Println("Application 'Journal-of-self' start")
	conConfig, err := config.LoadConfig("config.json")
	if err != nil {
		panic(err)
	}
	services, err := models.NewServices(conConfig.Database)
	defer services.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	u, err := models.CreateUser("a", "a", "a", "a")
	if err != nil {
		panic(err)
	}
	fmt.Println(services.Uservice.Insert(*u))
	fmt.Println(services)

	uc := controllers.NewUserController(services.Uservice)
	r := mux.NewRouter()
	r.HandleFunc("/users", uc.Signup).Methods("POST")
	http.ListenAndServe(":1117", r)
}
