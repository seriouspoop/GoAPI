package types

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	UserStore interface {
		GetUserByEmail(ctx context.Context, email string) (*User, error)
		GetUserByID(ctx context.Context, id string) (*User, error)
		CreateUser(ctx context.Context, user User) error
	}

	RegisterUserPayload struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	User struct {
		ID        primitive.ObjectID `bson:"_id, omitempty"`
		FirstName string             `bson:"first_name, omitempty"`
		LastName  string             `bson:"last_name, omitempty"`
		Email     string             `bson:"email, omitempty"`
		Password  string             `bson:"password, omitempty"`
		CreatedAt time.Time          `bson:"createdAt, omitempty"`
	}
)
