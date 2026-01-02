package lesson

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) Repository {
	return &MongoRepository{
		collection: db.Collection("lessons"),
	}
}

func (r *MongoRepository) Create(ctx context.Context, lesson *Lesson) error {
	now := time.Now()

	lesson.ID = primitive.NewObjectID()
	lesson.CreatedAt = now
	lesson.UpdatedAt = now

	_, err := r.collection.InsertOne(ctx, lesson)
	return err
}

func (r *MongoRepository) FindByCourse(ctx context.Context, courseID primitive.ObjectID) ([]Lesson, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"course_id": courseID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var lessons []Lesson
	if err := cursor.All(ctx, &lessons); err != nil {
		return nil, err
	}

	return lessons, nil
}
