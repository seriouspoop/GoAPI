package store

import (
	"context"

	"github.com/seriouspoop/GoAPI/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db *mongo.Database
}

func NewStore(db *mongo.Database) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	userColl := s.db.Collection("users")
	filter := bson.D{{Key: "email", Value: email}}

	user := new(types.User)
	err := userColl.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	return nil, nil
}

func (s *Store) CreateUser(ctx context.Context, user types.User) error {
	userColl := s.db.Collection("users")
	_, err := userColl.InsertOne(ctx, user)
	return err
}
