package mongodb_repository

import (
	"context"
	"fmt"
	"server/internal/app/repositories"
	"server/internal/app/repositories/mongodb_repository/mongo_configs"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type ApplicationsRepository struct {
	dbClient   *mongo.Client
	collection *mongo.Collection
	logger     *zap.Logger
	timeout    time.Duration
}

var _ repositories.ApplicationsRepository = (*ApplicationsRepository)(nil)

func NewApplicationsRepository(db *mongo.Client, l *zap.Logger, timeout time.Duration) *ApplicationsRepository {
	mongoDB := db.Database(mongo_configs.DBname)
	collection := mongoDB.Collection(mongo_configs.CollectionName)

	return &ApplicationsRepository{
		dbClient:   db,
		collection: collection,
		logger:     l,
		timeout:    timeout,
	}
}

func (r *ApplicationsRepository) GetApplicationNames() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	res, err := r.collection.Distinct(ctx, "appinfo.name", bson.D{})
	if err != nil {
		return nil, err
	}

	applications := make([]string, len(res))
	for i, v := range res {
		applications[i] = fmt.Sprint(v)
	}
	
	return applications, nil
}
