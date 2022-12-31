package main

import (
	"os"
	"testing"

	"github.com/abassGarane/muscles/domain"
)

var hd *handler

func TestMain(m *testing.M) {
	repo := initDB()
	serv := domain.NewService(repo)
	hd = NewHandler(serv)
	os.Exit(m.Run())
}
