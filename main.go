package main

import (
	"fmt"

	"github.com/daniel-z-johnson/journal-of-self/config"

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
}
