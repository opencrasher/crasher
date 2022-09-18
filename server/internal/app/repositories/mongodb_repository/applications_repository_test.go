package mongodb_repository

import (
	"errors"
	mock_repositories "server/internal/app/repositories/mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func TestNewApplicationsRepository(t *testing.T) {
	ctxTimeout := 5
	applicationsRepository := NewApplicationsRepository(&mongo.Client{}, &zap.Logger{}, ctxTimeout)

	require.Equal(t, "crasher", applicationsRepository.collection.Database().Name())
	require.Equal(t, "coredumps", applicationsRepository.collection.Name())
	require.Equal(t, &zap.Logger{}, applicationsRepository.logger)
	require.Equal(t, ctxTimeout, applicationsRepository.timeout)
}
func TestGetApplicationNames(t *testing.T) {
	t.Parallel()

	slowResponse := time.Second * 6

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
			slice: nil,
			error: nil,
		},
		{
			name: "get error while getting application names",
			stubs: func(r *mock_repositories.MockApplicationsRepository, slice []string) {
				r.EXPECT().GetApplicationNames().Return(slice, errors.New("error"))
			},
			slice: nil,
			error: errors.New("error"),
		},
		{
			name: "get timeout error while getting application names",
			stubs: func(r *mock_repositories.MockApplicationsRepository, slice []string) {
				r.EXPECT().GetApplicationNames().DoAndReturn(func() ([]string, error) {
					time.Sleep(slowResponse)
					return slice, errors.New("error")
				})
			},
			slice: nil,
			error: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.stubs(r, test.slice)

			result, err := r.GetApplicationNames()
			require.Equal(t, test.error, err)
			require.Equal(t, len(test.slice), len(result))
		})
	}
}
