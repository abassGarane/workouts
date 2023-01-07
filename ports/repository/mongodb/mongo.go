package mongodb

import (
	"context"
	"log"
	"time"

	"github.com/abassGarane/muscles/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	// client   *mongo.Client
	timeout  time.Duration
	database string
	ctx      context.Context
	col      *mongo.Collection
	usersCol *mongo.Collection
}

func newClient(mongoUrl string, ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		log.Print("Unable to disconnect client", err)
	// 	}
	// }()
	return client, nil
}
func NewMongoRepository(mongoURL, database string, timeout int, ctx context.Context) (domain.Repository, error) {
	client, err := newClient(mongoURL, ctx)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	repo := &mongoRepository{
		database: database,
		col:      client.Database("muscles").Collection("workouts"),
		ctx:      ctx,
		timeout:  time.Duration(timeout) * time.Second,
		usersCol: client.Database("muscles").Collection("users"),
	}
	repo.usersCol.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.M{"email": 1},
			Options: options.Index().SetUnique(true),
		},
	)
	return repo, nil
}
