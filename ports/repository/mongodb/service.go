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
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *mongoRepository) AddWorkout(workout *domain.Workout) (*domain.Workout, error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	defer cancel()
	if workout.ID = primitive.NewObjectID(); false {
		return nil, errors.New("Unable to created new objectid")
	}
	_, err := m.col.InsertOne(ctx, &workout)
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Repository.addWorkout")
	}
	return workout, nil
}

func (m *mongoRepository) GetWorkout(id string) (*domain.Workout, error) {
	workout := domain.Workout{}
	if primitive.IsValidObjectID(id) {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			if err == primitive.ErrInvalidHex {
				return nil, errors.Wrap(err, "Invalid id type")
			} else if objID == primitive.NilObjectID || !primitive.IsValidObjectID(id) {
				return nil, errors.New("Unable to create new objectid")
			}
		}
		if objID.IsZero() {
			return nil, errors.New("id can not be empty :: repo.mongoRepository.GetWorkout")
		}
		fmt.Println(objID, err)
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
		defer cancel()
		fmt.Printf("mongo result :: %#v", m.col.FindOne(ctx, bson.M{"_id": objID}))
		res := m.col.FindOne(ctx, bson.M{"_id": objID})
		if res.Err() == mongo.ErrNoDocuments {
			return nil, errors.Wrap(err, "No documents with that id in the repo")
		}
		fmt.Printf("Retrieved workout %#v\n", workout)
		err = res.Decode(&workout)
		if err != nil {
			return nil, errors.Wrap(err, "repo.mongoRepository.GetWorkout")
		}
		log.Println(workout)
		return &workout, nil

	}
	// check if objectID is empty

	return nil, errors.New("Invalid hex type")
}

func (m *mongoRepository) GetWorkouts() ([]*domain.Workout, error) {
	workouts := []*domain.Workout{}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	defer cancel()
	opts := options.Find().SetSort(bson.M{"created_at": -1})
	cursor, err := m.col.Find(ctx, bson.M{}, opts)
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
	if err != nil {
		return errors.Wrap(err, "repo.DeleteWorkout")
	}
	existing := m.col.FindOne(context.Background(), bson.M{"_id": objId})
	if existing == nil {
		return errors.New("Document is not in the database")
	}
	res := m.col.FindOneAndDelete(context.Background(), bson.M{"_id": objId})
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
		return nil, errors.Wrap(err, "repo.UpdateWorkout")
	}
	// //retrieve old workout
	// if err = col.FindOne(context.Background(), bson.M{"_id":objId}).Decode(&oldWorkout); err != nil{
	// 	//TODO : Save workout as a new workout
	// 	return nil, errors.Wrap(err, "Workout doesnt exist")
	// }
	_, err = m.col.ReplaceOne(context.Background(), bson.M{"_id": objId}, workout)
	// replace old workout with new workout
	if err != nil {
		return nil, errors.Wrap(err, "Workout doesnt exist")
	}
	updatedWorkout = workout
	updatedWorkout.ID = objId

	fmt.Println("updated object", updatedWorkout)
	return updatedWorkout, nil
}
