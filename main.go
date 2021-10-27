package main

import (
	"flag"
	"fmt"
	"github.com/daniel-z-johnson/journal-of-self/views"
	"net/http"

	"github.com/daniel-z-johnson/journal-of-self/config"
	"github.com/daniel-z-johnson/journal-of-self/controllers"
	"github.com/go-chi/chi/v5"

	"github.com/daniel-z-johnson/journal-of-self/models"
)

func main() {
	migrate := flag.Bool("migrate", false, "If set, will do migration")
	flag.Parse()

	fmt.Println("Application 'Journal-of-self' start")
	conConfig, err := config.LoadConfig("config.json")
	if err != nil {
		panic(err)
	}
	services, err := models.NewServices(conConfig.Database, *migrate)
	defer services.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	u, err := models.CreateUser("a", "a", "a", "a")
	if err != nil {
		panic(err)
	}
	fmt.Println(services.UserService.Insert(*u))
	fmt.Println(services)

	uc := controllers.NewUserController(services.UserService)
	r := chi.NewRouter()
	r.Get("/users", uc.Signup)
	r.Handle("/static/*", views.StaticFilesServer())
	http.ListenAndServe(":1117", r)
}
