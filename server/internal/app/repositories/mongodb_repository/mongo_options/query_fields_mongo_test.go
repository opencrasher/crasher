package mongo_options

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestSetApplication(t *testing.T) {
	filter := bson.M{}
	require.Equal(t, 0, len(filter))

	options := options.Find()
	setter := SetApplication("app")
	setter(filter, options)
	res := filter["appinfo.name"]

	require.Equal(t, "app", res)
	require.Equal(t, true, len(filter) > 0)
}

func TestSetStartTimestamp(t *testing.T) {
	filter := bson.M{}
	require.Equal(t, 0, len(filter))

	requiredDate := primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -1))

	options := options.Find()
	setter := SetStartTimestamp(requiredDate)
	setter(filter, options)
	res := filter["timestamp"]

	require.Equal(t, primitive.M(primitive.M{"$gte": requiredDate}), res)
	require.Equal(t, true, len(filter) > 0)
}

func TestSetEndTimestamp(t *testing.T) {
	filter := bson.M{}
	require.Equal(t, 0, len(filter))

	requiredDate := primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, 0))

	options := options.Find()
	setter := SetEndTimestamp(requiredDate)
	setter(filter, options)
	res := filter["timestamp"]

	require.Equal(t, primitive.M(primitive.M{"$lt": requiredDate}), res)
	require.Equal(t, true, len(filter) > 0)
}

func TestSetLimit(t *testing.T) {
	filter := bson.M{}
	require.Equal(t, 0, len(filter))

	options := options.Find()
	require.Equal(t, true, options.Limit == nil)

	setter := SetLimit(1)
	setter(filter, options)

	require.Equal(t, int64(1), *options.Limit)
	require.Equal(t, false, options == nil)
}

func TestSetOffeset(t *testing.T) {
	filter := bson.M{}
	require.Equal(t, 0, len(filter))

	options := options.Find()
	require.Equal(t, true, options.Skip == nil)

	setter := SetOffset(1)
	setter(filter, options)

	require.Equal(t, int64(1), *options.Skip)
	require.Equal(t, false, options == nil)
}
