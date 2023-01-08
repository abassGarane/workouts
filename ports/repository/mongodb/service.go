package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/abassGarane/muscles/domain"
	"github.com/abassGarane/muscles/domain/models"
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
	_, err := m.workoutsCol.InsertOne(ctx, &workout)
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
		res := m.workoutsCol.FindOne(ctx, bson.M{"_id": objID})
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
	cursor, err := m.workoutsCol.Find(ctx, bson.M{}, opts)
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
	existing := m.workoutsCol.FindOne(context.Background(), bson.M{"_id": objId})
	if existing == nil {
		return errors.New("Document is not in the database")
	}
	res := m.workoutsCol.FindOneAndDelete(context.Background(), bson.M{"_id": objId})
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
	_, err = m.workoutsCol.ReplaceOne(context.Background(), bson.M{"_id": objId}, workout)
	// replace old workout with new workout
	if err != nil {
		return nil, errors.Wrap(err, "Workout doesnt exist")
	}
	updatedWorkout = workout
	updatedWorkout.ID = objId

	fmt.Println("updated object", updatedWorkout)
	return updatedWorkout, nil
}

// users services
func (m *mongoRepository) CreateUser(user *models.User) (*models.User, error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	defer cancel()
	if user.ID = primitive.NewObjectID(); false {
		return nil, errors.New("Unable to created new objectid")
	}
	//check if user exists
	res := m.usersCol.FindOne(ctx, bson.M{"email": user.Email})
	if res.Err() == nil {
		return nil, errors.Wrap(res.Err(), "repo.Create user"+"User exists already")
	}
	_, err := m.usersCol.InsertOne(ctx, &user)
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Repository.addWorkout")
	}
	return user, nil

}
func (m *mongoRepository) UpdateUser(email string, user *models.User) (*models.User, error) {
	var updatedUser *models.User

	_, err := m.usersCol.ReplaceOne(context.Background(), bson.M{"email": email}, user)
	// replace old workout with new workout
	if err != nil {
		return nil, errors.Wrap(err, "Could not update user")
	}
	updatedUser = user
	return updatedUser, nil

}
func (m *mongoRepository) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	defer cancel()
	if user.ID = primitive.NewObjectID(); false {
		return nil, errors.New("Unable to created new objectid")
	}
	//check if user exists
	res := m.usersCol.FindOne(ctx, bson.M{"email": email})
	if err := res.Decode(&user); err != nil {
		return nil, errors.Wrap(res.Err(), "repo.GetUserByEmail user"+"Could not find user")
	}
	return user, nil
}
