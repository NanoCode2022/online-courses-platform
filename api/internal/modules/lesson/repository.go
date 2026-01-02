package lesson

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	Create(ctx context.Context, lesson *Lesson) error
	FindByCourse(ctx context.Context, courseID primitive.ObjectID) ([]Lesson, error)
}
