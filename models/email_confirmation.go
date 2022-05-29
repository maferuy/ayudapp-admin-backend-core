package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailConfirmation struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Confirmed bool               `bson:"status" json:"status"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
