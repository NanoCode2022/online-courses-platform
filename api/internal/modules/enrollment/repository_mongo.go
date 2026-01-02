package enrollment

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
		collection: db.Collection("enrollments"),
	}
}

func (r *MongoRepository) Enroll(ctx context.Context, userID string, courseID primitive.ObjectID) error {
	enrollment := Enrollment{
		UserID:    userID,
		CourseID:  courseID,
		CreatedAt: time.Now(),
	}

	_, err := r.collection.InsertOne(ctx, enrollment)
	return err
}

func (r *MongoRepository) IsEnrolled(ctx context.Context, userID string, courseID primitive.ObjectID) (bool, error) {
	count, err := r.collection.CountDocuments(ctx, bson.M{
		"user_id":   userID,
		"course_id": courseID,
	})
	return count > 0, err
}
