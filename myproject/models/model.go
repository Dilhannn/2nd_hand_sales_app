// models.go

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Email    string             `json:"email"`
	Basket   []string           `json:"basket"`
}

type Photo struct {
	ID          primitive.ObjectID `json:"id"`
	UserID      string             `json:"userId"`
	URL         string             `json:"url"`
	Tags        string             `json:"tags"`
	Price       string             `json:"price"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"createdAt"`
}

type PhotoCreate struct {
	UserID      string    `json:"userId"`
	URL         string    `json:"url"`
	Tags        string    `json:"tags"`
	Price       string    `json:"price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Comment struct {
	UserID    string    `json:"userId"`
	Content   string    `json:"content"`
	Timestamp string    `json:"timestamp"`
	CreatedAt time.Time `json:"createdAt"`
}

type Cart struct {
	ID     primitive.ObjectID `json:"id"`
	UserID string             `json:"userId"`
	Photo  Photo              `json:"photo"`
}

type Search struct {
	Field []string `json:"field"`
}

type IDName struct {
	ID primitive.ObjectID `json:"id"`
}

type IDUserName struct {
	UserID string `json:"userId"`
}

type Product struct {
	ID        primitive.ObjectID `json:"id"`
	UserID    string             `json:"userId"`
	CreatedAt time.Time          `json:"createdAt"`
}
