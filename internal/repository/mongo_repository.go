package repository

import (
	"context"
	"time"

	"go-data-ingestor/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	collection *mongo.Collection
}

func NewMongoRepo(uri, db, col string) (*MongoRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	coll := client.Database(db).Collection(col)
	return &MongoRepo{collection: coll}, nil
}

func (r *MongoRepo) Save(posts []models.Post) error {
	docs := make([]interface{}, len(posts))
	for i, post := range posts {
		docs[i] = post
	}
	_, err := r.collection.InsertMany(context.Background(), docs)
	return err
}
