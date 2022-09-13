package mongodb_repository

import (
	"context"
	"fmt"
	"server/internal/app/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type ApplicationsRepository struct {
	dbClient   *mongo.Client
	collection *mongo.Collection
	logger     *zap.Logger
	timeout    int
}

var _ repositories.ApplicationsRepository = (*ApplicationsRepository)(nil)

func NewApplicationsRepository(db *mongo.Client, l *zap.Logger) *ApplicationsRepository {
	mongoDB := db.Database("crasher")
	collection := mongoDB.Collection("coredumps")
	timeout, err := ParseCtxTimeoutEnv()
	if err != nil {
		l.Fatal(
			"failed to parse ctx timeout env",
			zap.Error(err),
		)
	}
	return &ApplicationsRepository{
		dbClient:   db,
		collection: collection,
		logger:     l,
		timeout:    timeout,
	}
}

func (r *ApplicationsRepository) GetApplicationNames() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.timeout)*time.Second)
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
