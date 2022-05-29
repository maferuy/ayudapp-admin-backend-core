package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty", json:"_id,omitempty"`
	Date      time.Time          `bson:"date", json:"date"`
	Duration  time.Duration      `bson:"duration", json:"duration"`
	Address   string             `bson:"address", json:"address"`
	Status    string             `bson:"status", json:"status"`
	CreatedBy primitive.ObjectID `bson:"created_by,omitempty", json:"created_by,omitempty"`
	Helper    primitive.ObjectID `bson:"helper,omitempty", json:"helper,omitempty"`
	CreatedAt time.Time          `bson:"created_at", json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at", json:"updated_at"`
}
