package course

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	Create(ctx context.Context, course *Course) error
	FindAll(ctx context.Context) ([]Course, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*Course, error)
}
