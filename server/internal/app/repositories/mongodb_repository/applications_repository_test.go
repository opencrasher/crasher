package mongodb_repository

import (
	"errors"
	mock_repositories "server/internal/app/repositories/mock"
	"server/internal/app/repositories/mongodb_repository/mongo_configs"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func TestNewApplicationsRepository(t *testing.T) {
	ctxTimeout := time.Duration(5) * time.Second
	applicationsRepository := NewApplicationsRepository(&mongo.Client{}, &zap.Logger{}, ctxTimeout)

	require.Equal(t, mongo_configs.DBname, applicationsRepository.collection.Database().Name())
	require.Equal(t, mongo_configs.CollectionName, applicationsRepository.collection.Name())
	require.NotNil(t, applicationsRepository.logger)
	require.Equal(t, ctxTimeout, applicationsRepository.timeout)
}
func TestGetApplicationNames(t *testing.T) {
	t.Parallel()

	c := gomock.NewController(t)
	defer c.Finish()
	r := mock_repositories.NewMockApplicationsRepository(c)

	tests := []struct {
		name  string
		stubs func(store *mock_repositories.MockApplicationsRepository, slice []string)
		slice []string
		error error
	}{
		{
			name: "get many application names",
			stubs: func(r *mock_repositories.MockApplicationsRepository, slice []string) {
				r.EXPECT().GetApplicationNames().Return(slice, nil)
			},
			slice: []string{"app1", "app2", "app3"},
			error: nil,
		},
		{
			name: "get one application name",
			stubs: func(r *mock_repositories.MockApplicationsRepository, slice []string) {
				r.EXPECT().GetApplicationNames().Return(slice, nil)
			},
			slice: []string{"app1"},
			error: nil,
		},
		{
			name: "get empty application names",
			stubs: func(r *mock_repositories.MockApplicationsRepository, slice []string) {
				r.EXPECT().GetApplicationNames().Return(slice, nil)
			},
			slice: []string{},
			error: nil,
		},
		{
			name: "get error while getting application names",
			stubs: func(r *mock_repositories.MockApplicationsRepository, slice []string) {
				r.EXPECT().GetApplicationNames().Return(slice, errors.New("error"))
			},
			slice: []string{},
			error: errors.New("error"),
		},
		{
			name: "get timeout error while getting application names",
			stubs: func(r *mock_repositories.MockApplicationsRepository, slice []string) {
				r.EXPECT().GetApplicationNames().DoAndReturn(func() ([]string, error) {
					time.Sleep(6 * time.Second)
					return slice, errors.New("error")
				})
			},
			slice: []string{},
			error: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.stubs(r, test.slice)

			result, err := r.GetApplicationNames()
			assert.Equal(t, test.error, err)
			assert.Equal(t, len(test.slice), len(result))
		})
	}
}
