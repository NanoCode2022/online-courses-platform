package lesson

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Lesson struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CourseID  primitive.ObjectID `bson:"course_id" json:"course_id"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	VideoURL  string             `bson:"video_url" json:"video_url"`
	Order     int                `bson:"order" json:"order"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
