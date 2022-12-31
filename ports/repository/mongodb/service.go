package mongodb

import (
	"context"
	"log"
	"time"

	"github.com/abassGarane/muscles/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *mongoRepository) AddWorkout(workout *domain.Workout) (*domain.Workout, error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	defer cancel()
	coll := m.client.Database("muscles").Collection("workouts")
	workout.ID = primitive.NewObjectID()
	_, err := coll.InsertOne(ctx, &workout)
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Repository.addWorkout")
	}
	return workout, nil
}

func (m *mongoRepository) GetWorkout(id string) (*domain.Workout, error) {
	workout := &domain.Workout{}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "repo.mongoRepository.GetWorkout")
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	defer cancel()
	coll := m.client.Database("muscles").Collection("workouts")
	err = coll.FindOne(ctx, bson.M{"_id": objID}).Decode(&workout)
	if err != nil {
		return nil, errors.Wrap(err, "repo.mongoRepository.GetWorkout")
	}
	log.Println(workout)
	return workout, nil
}
