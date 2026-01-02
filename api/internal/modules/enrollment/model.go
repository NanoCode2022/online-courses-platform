package enrollment

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Enrollment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    string             `bson:"user_id" json:"user_id"`
	CourseID  primitive.ObjectID `bson:"course_id" json:"course_id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
