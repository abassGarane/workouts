package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/abassGarane/muscles/domain"
	"github.com/abassGarane/muscles/ports/repository/mongodb"
)

var (
	mongoURL     string
	mongoDB      string
	mongoTimeout int
	repo         domain.Repository
)

func initDB() domain.Repository {
	var err error
	flag.StringVar(&mongoURL, "MONGO_URL", "mongodb://root:root@localhost:27017", "Mongodb connection string")
	flag.IntVar(&mongoTimeout, "MONGO_TIMEOUT", 10, "Mongo timeout")
	flag.StringVar(&mongoDB, "MONGO_DB", "muscles", "Mongo db string")
	//Repo
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	defer cancel()
	repo, err = mongodb.NewMongoRepository(mongoURL, mongoDB, mongoTimeout, ctx)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}
