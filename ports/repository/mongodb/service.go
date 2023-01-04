package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/abassGarane/muscles/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoRepository) AddWorkout(workout *domain.Workout) (*domain.Workout, error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	defer cancel()
	coll := m.client.Database("muscles").Collection("workouts")
	if workout.ID = primitive.NewObjectID().Hex(); false {
		return nil, errors.New("Unable to created new objectid")
	}
	_, err := coll.InsertOne(ctx, &workout)
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Repository.addWorkout")
	}
	return workout, nil
}

func (m *mongoRepository) GetWorkout(id string) (*domain.Workout, error) {
	workout := &domain.Workout{}
	if id == "" {
		return nil, errors.New("id is empty")
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		if err == primitive.ErrInvalidHex {
			return nil, errors.Wrap(err, "Invalid id type")
		}
		if objID == primitive.NilObjectID || !primitive.IsValidObjectID(id) {
			return nil, errors.New("Unable to created new objectid")
		}
	}
	fmt.Println(objID, err)
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

func (m *mongoRepository) GetWorkouts() ([]*domain.Workout, error) {
	workouts := []*domain.Workout{}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	defer cancel()
	coll := m.client.Database("muscles").Collection("workouts")
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(err, "repo.mongoRepository.GetWorkout :: No document found")
		} else {
			return nil, errors.Wrap(err, "repo.mongoRepository.GetWorkout")
		}
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &workouts); err != nil {
		return nil, errors.Wrap(err, "repo.mongoRepository.GetWorkout")
	}
	return workouts, nil
}

func (m *mongoRepository) DeleteWorkout(id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	fmt.Println(objId)
	if err != nil {
		return errors.Wrap(err, "repo.DeleteWorkout")
	}
	col := m.client.Database("muscles").Collection("workouts")
	existing := col.FindOne(context.Background(), bson.M{"_id": objId})
	if existing == nil {
		return errors.New("Document is not in the database")
	}
	res := col.FindOneAndDelete(context.Background(), bson.M{"_id": objId})
	fmt.Println(res)
	if res.Err() != nil {
		return errors.New("Unable to delete workout")
	}
	return nil
}

func (m *mongoRepository) UpdateWorkout(id string, workout *domain.Workout) (*domain.Workout, error) {
	var updatedWorkout *domain.Workout

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "repo.DeleteWorkout")
	}
	col := m.client.Database("muscles").Collection("workouts")
	// //retrieve old workout
	// if err = col.FindOne(context.Background(), bson.M{"_id":objId}).Decode(&oldWorkout); err != nil{
	// 	//TODO : Save workout as a new workout
	// 	return nil, errors.Wrap(err, "Workout doesnt exist")
	// }
	_, err = col.ReplaceOne(context.Background(), bson.M{"_id": objId}, workout)
	// replace old workout with new workout
	if err != nil {
		return nil, errors.Wrap(err, "Workout doesnt exist")
	}
	updatedWorkout = workout
	updatedWorkout.ID = objId.Hex()

	fmt.Println("updated object", updatedWorkout)
	return updatedWorkout, nil
}
