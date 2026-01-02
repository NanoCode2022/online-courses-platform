package enrollment

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	Enroll(ctx context.Context, userID string, courseID primitive.ObjectID) error
	IsEnrolled(ctx context.Context, userID string, courseID primitive.ObjectID) (bool, error)
}
