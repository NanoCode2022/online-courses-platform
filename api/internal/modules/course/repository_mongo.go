package course

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
		collection: db.Collection("courses"),
	}
}

func (r *MongoRepository) Create(ctx context.Context, course *Course) error {
	now := time.Now()

	course.ID = primitive.NewObjectID()
	course.CreatedAt = now
	course.UpdatedAt = now

	_, err := r.collection.InsertOne(ctx, course)
	return err
}

func (r *MongoRepository) FindAll(ctx context.Context) ([]Course, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var courses []Course
	if err := cursor.All(ctx, &courses); err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *MongoRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*Course, error) {
	var course Course

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&course)
	if err != nil {
		return nil, err
	}

	return &course, nil
}
