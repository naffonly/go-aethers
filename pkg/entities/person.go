package entities

import (
	"time"
	"uuid"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Author    string             `json:"author" bson:"author,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type PersonDB struct {
	PersonID  uuid.UUID `gorm:"type:uuid;primaryKey"`
	LastName  string    `gorm:"type:varchar(255);not null"`
	FirstName string    `gorm:"type:varchar(255)"`
	City      string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
