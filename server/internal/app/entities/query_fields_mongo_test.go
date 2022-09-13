package entities

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Test_SetApplication(t *testing.T) {
	filter := bson.M{}
	options := options.Find()
	setter := SetApplication("app")
	setter(filter, options)
	res := filter["appinfo.name"]
	require.Equal(t, "app", res)
}
