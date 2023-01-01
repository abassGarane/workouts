package main

import (
	"log"
	"net/http"

	"github.com/abassGarane/muscles/domain"
)

var (
	service domain.Service
)

func main() {
	env := initEnv()
	PORT:= env["PORT"]
	repo := initDB()
	service = domain.NewService(repo)
	router := initRouter(service)
	if err := http.ListenAndServe(PORT.(string), router); err != nil {
		log.Fatal(err)
	}
}
