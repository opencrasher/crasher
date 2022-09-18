package mongodb_repository

import (
	"context"
	"server/internal/app/entities"
	"server/internal/app/repositories"
	"server/internal/app/repositories/mongodb_repository/mongo_options"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type CoreDumpsRepository struct {
	dbClient   *mongo.Client
	collection *mongo.Collection
	logger     *zap.Logger
	timeout    int
}

var _ repositories.CoreDumpsRepository = (*CoreDumpsRepository)(nil)

func NewCoreDumpsRepository(db *mongo.Client, l *zap.Logger, timeout int) *CoreDumpsRepository {
	mongoDB := db.Database("crasher")
	collection := mongoDB.Collection("coredumps")

	return &CoreDumpsRepository{
		dbClient:   db,
		collection: collection,
		logger:     l,
		timeout:    timeout,
	}
}

func (r *CoreDumpsRepository) GetCoreDumps(parameters ...interface{}) ([]entities.CoreDump, error) {
	setters := make([]mongo_options.OptionsMongo, len(parameters))
	for _, setter := range parameters {
		if isMongo, ok := setter.(mongo_options.OptionsMongo); ok {
			setters = append(setters, isMongo)
		}
	}

	options := options.Find()
	filter := bson.M{}
	for _, setter := range setters {
		setter(filter, options)
	}

	var result []entities.CoreDump

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.timeout)*time.Second)
	defer cancel()

	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &result)
	return result, err
}

func (r *CoreDumpsRepository) GetCoreDumpByID(id string) (entities.CoreDump, error) {
	var coreDump entities.CoreDump

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.timeout)*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.D{{Key: "id", Value: id}}).Decode(&coreDump)
	return coreDump, err
}

func (r *CoreDumpsRepository) UpdateCoreDump(parameters interface{}) error {
	return nil
}

func (r *CoreDumpsRepository) AddCoreDump(coreDump entities.CoreDump) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.timeout)*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, coreDump)
	return err
}

func (r *CoreDumpsRepository) DeleteCoreDump(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.timeout)*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.D{{Key: "id", Value: id}})
	return err
}

func (r *CoreDumpsRepository) DeleteAllCoreDumps() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.timeout)*time.Second)
	defer cancel()

	_, err := r.collection.DeleteMany(ctx, bson.D{})
	return err
}
