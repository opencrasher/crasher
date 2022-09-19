package mongo_options

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestSetApplication(t *testing.T) {
	filter := bson.M{}
	assert.Equal(t, 0, len(filter))

	options := options.Find()
	setter := SetApplication("app")
	setter(filter, options)
	res := filter["appinfo.name"]

	assert.Equal(t, "app", res)
}

func TestSetStartTimestamp(t *testing.T) {
	filter := bson.M{}
	assert.Equal(t, 0, len(filter))

	requiredDate := primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -1))

	options := options.Find()
	setter := SetStartTimestamp(requiredDate)
	setter(filter, options)
	res := filter["timestamp"]

	assert.Equal(t, primitive.M(primitive.M{"$gte": requiredDate}), res)
}

func TestSetEndTimestamp(t *testing.T) {
	filter := bson.M{}
	assert.Equal(t, 0, len(filter))

	requiredDate := primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, 0))

	options := options.Find()
	setter := SetEndTimestamp(requiredDate)
	setter(filter, options)
	res := filter["timestamp"]

	assert.Equal(t, primitive.M(primitive.M{"$lt": requiredDate}), res)
}

func TestSetLimit(t *testing.T) {
	filter := bson.M{}
	assert.Equal(t, 0, len(filter))

	options := options.Find()
	assert.Equal(t, true, options.Limit == nil)

	setter := SetLimit(1)
	setter(filter, options)

	assert.Equal(t, int64(1), *options.Limit)
}

func TestSetOffeset(t *testing.T) {
	filter := bson.M{}
	assert.Equal(t, 0, len(filter))

	options := options.Find()
	assert.Equal(t, true, options.Skip == nil)

	setter := SetOffset(1)
	setter(filter, options)

	assert.Equal(t, int64(1), *options.Skip)
}
