package repository

import (
	"albo/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	db  *mongo.Database
	ctx context.Context
}

func NewRepository(db *mongo.Database, ctx context.Context) *repository {
	return &repository{
		db:  db,
		ctx: ctx,
	}
}

func (repo *repository) Get(character string) (*domain.Colaborator, error) {
	filter := bson.D{}
	options := options.FindOne().SetSort(bson.D{{"_id", -1}})

	var colaborator domain.Colaborator
	if err := repo.db.Collection("colaborators").
		FindOne(context.Background(), filter, options).Decode(&colaborator); err != nil {
		return nil, err
	}

	return &colaborator, nil
}

func (repo *repository) Save(data *domain.Colaborator) error {
	_, err := repo.db.Collection("colaborators").InsertOne(context.Background(), data)
	if err != nil {
		return err
	}

	return nil
}
