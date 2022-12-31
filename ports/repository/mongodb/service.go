package mongodb

import (
	"context"
	"time"

	"github.com/abassGarane/muscles/domain"
	"github.com/pkg/errors"
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
