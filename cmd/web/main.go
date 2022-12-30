package main

import (
	"log"
	"net/http"

	"github.com/abassGarane/muscles/domain"
)

func main() {
	env := initEnv()
	PORT, _ := env["PORT"]

	router := initRouter()
	repo := initDB()
	service := domain.NewService(repo)
	if err := http.ListenAndServe(PORT.(string), router); err != nil {
		log.Fatal(err)
	}
}
