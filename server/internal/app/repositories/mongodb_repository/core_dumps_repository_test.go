package mongodb_repository

import (
	"errors"
	"fmt"
	"server/internal/app/entities"
	mock_repositories "server/internal/app/repositories/mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func TestNewCoreDumpsRepository(t *testing.T) {
	t.Parallel()

	ctxTimeout := 5

	applicationsRepository := NewApplicationsRepository(&mongo.Client{}, &zap.Logger{}, ctxTimeout)

	require.Equal(t, "crasher", applicationsRepository.collection.Database().Name())
	require.Equal(t, "coredumps", applicationsRepository.collection.Name())
	require.Equal(t, &zap.Logger{}, applicationsRepository.logger)
	require.Equal(t, ctxTimeout, applicationsRepository.timeout)
}
func TestAddCoreDump(t *testing.T) {
	t.Parallel()

	slowResponse := time.Second * 6

	c := gomock.NewController(t)
	defer c.Finish()
	r := mock_repositories.NewMockCoreDumpsRepository(c)

	randomDump := generateRandomSliceOfCoreDumps(1)
	tests := []struct {
		name  string
		dump  entities.CoreDump
		stubs func(store *mock_repositories.MockCoreDumpsRepository, dump entities.CoreDump)
		error error
	}{
		{
			name: "add dump",
			dump: randomDump[0],
			stubs: func(r *mock_repositories.MockCoreDumpsRepository, dump entities.CoreDump) {
				r.EXPECT().AddCoreDump(dump).Return(nil)
			},
			error: nil,
		},
		{
			name: "add dump error",
			dump: entities.CoreDump{},
			stubs: func(r *mock_repositories.MockCoreDumpsRepository, dump entities.CoreDump) {
				r.EXPECT().AddCoreDump(dump).Return(errors.New("error"))

			},
			error: errors.New("error"),
		},
		{
			name: "add dump timeout error",
			dump: randomDump[0],
			stubs: func(r *mock_repositories.MockCoreDumpsRepository, dump entities.CoreDump) {
				r.EXPECT().AddCoreDump(dump).DoAndReturn(func(dump entities.CoreDump) error {
					time.Sleep(slowResponse)
					return errors.New("error")
				})
			},
			error: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.stubs(r, test.dump)

			err := r.AddCoreDump(test.dump)
			require.Equal(t, test.error, err)
		})
	}
}

func TestGetCoreDumps(t *testing.T) {
	t.Parallel()

	slowResponse := time.Second * 6

	c := gomock.NewController(t)
	defer c.Finish()
	r := mock_repositories.NewMockCoreDumpsRepository(c)

	sliceOfDumps := generateRandomSliceOfCoreDumps(5)

	tests := []struct {
		name  string
		stubs func(store *mock_repositories.MockCoreDumpsRepository)
		dumps []entities.CoreDump
		error error
	}{
		{
			name: "get core dumps",
			stubs: func(r *mock_repositories.MockCoreDumpsRepository) {
				r.EXPECT().GetCoreDumps().Return(sliceOfDumps, nil)
			},
			dumps: sliceOfDumps,
			error: nil,
		},
		{
			name: "get core dumps error",
			stubs: func(r *mock_repositories.MockCoreDumpsRepository) {
				r.EXPECT().GetCoreDumps().Return(nil, errors.New("error"))
			},
			dumps: nil,
			error: errors.New("error"),
		},
		{
			name: "get core dumps error timeout",
			stubs: func(r *mock_repositories.MockCoreDumpsRepository) {
				r.EXPECT().GetCoreDumps().Do(func(options ...interface{}) {
					time.Sleep(slowResponse)
				}).Return(nil, errors.New("error"))
			},
			dumps: nil,
			error: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.stubs(r)

			res, err := r.GetCoreDumps()
			require.Equal(t, test.error, err)
			require.Equal(t, test.dumps, res)
		})
	}
}

func TestGetCoreDumpByID(t *testing.T) {
	t.Parallel()

	slowResponse := time.Second * 6

	c := gomock.NewController(t)
	defer c.Finish()
	r := mock_repositories.NewMockCoreDumpsRepository(c)

	sliceOfDumps := generateRandomSliceOfCoreDumps(1)

	tests := []struct {
		name   string
		stubs  func(store *mock_repositories.MockCoreDumpsRepository, id string)
		dumpID string
		error  error
	}{
		{
			name: "get core dump by id",
			stubs: func(r *mock_repositories.MockCoreDumpsRepository, id string) {
				r.EXPECT().GetCoreDumpByID(id).Return(sliceOfDumps[0], nil)
			},
			dumpID: sliceOfDumps[0].ID,
			error:  nil,
		},
		{
			name: "get core dump by id error",
			stubs: func(r *mock_repositories.MockCoreDumpsRepository, id string) {
				r.EXPECT().GetCoreDumpByID(id).Return(entities.CoreDump{}, errors.New("error"))
			},
			dumpID: "",
			error:  errors.New("error"),
		},
		{
			name: "get core dump by id error timeout",
			stubs: func(r *mock_repositories.MockCoreDumpsRepository, id string) {
				r.EXPECT().GetCoreDumpByID(id).Do(func(id string) {
					time.Sleep(slowResponse)
				}).Return(entities.CoreDump{}, errors.New("error"))
			},
			dumpID: "",
			error:  errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.stubs(r, test.dumpID)

			res, err := r.GetCoreDumpByID(test.dumpID)
			require.Equal(t, test.error, err)
			require.Equal(t, test.dumpID, res.ID)
		})
	}
}

func TestDeleteCoreDump(t *testing.T) {
	t.Parallel()

	slowResponse := time.Second * 6

	c := gomock.NewController(t)
	defer c.Finish()
	r := mock_repositories.NewMockCoreDumpsRepository(c)

	tests := []struct {
		name   string
		stubs  func(store *mock_repositories.MockCoreDumpsRepository, dumpID string)
		dumpID string
		error  error
	}{
		{
			name: "delete dump",
			stubs: func(r *mock_repositories.MockCoreDumpsRepository, dumpID string) {
				r.EXPECT().DeleteCoreDump(dumpID).Return(nil)
			},
			dumpID: "dsfdfds454",
			error:  nil,
		},
		{
			name: "delete dump error empty id",
			stubs: func(r *mock_repositories.MockCoreDumpsRepository, dumpID string) {
				r.EXPECT().DeleteCoreDump(dumpID).Return(errors.New("error"))
			},
			dumpID: "",
			error:  errors.New("error"),
		},
		{
			name: "delete dump timeout error",
			stubs: func(r *mock_repositories.MockCoreDumpsRepository, dumpID string) {
				r.EXPECT().DeleteCoreDump(dumpID).Do(func(dumpID string) {
					time.Sleep(slowResponse)
				}).Return(errors.New("error"))
			},
			dumpID: "",
			error:  errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.stubs(r, test.dumpID)

			err := r.DeleteCoreDump(test.dumpID)
			require.Equal(t, test.error, err)
		})
	}
}

func TestDeleteAllCoreDump(t *testing.T) {
	t.Parallel()

	slowResponse := time.Second * 6

	c := gomock.NewController(t)
	defer c.Finish()

	r := mock_repositories.NewMockCoreDumpsRepository(c)

	tests := []struct {
		name  string
		stubs func(store *mock_repositories.MockCoreDumpsRepository)
		error error
	}{
		{
			name: "delete all dumps",
			stubs: func(r *mock_repositories.MockCoreDumpsRepository) {
				r.EXPECT().DeleteAllCoreDumps().Return(nil)
			},
			error: nil,
		},
		{
			name: "delete dump timeout error",
			stubs: func(r *mock_repositories.MockCoreDumpsRepository) {
				r.EXPECT().DeleteAllCoreDumps().Do(func() {
					time.Sleep(slowResponse)
				}).Return(errors.New("error"))
			},
			error: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.stubs(r)

			err := r.DeleteAllCoreDumps()
			require.Equal(t, test.error, err)
		})
	}
}

func generateRandomSliceOfCoreDumps(quantity int) []entities.CoreDump {
	type osInfo struct {
		Name    string
		Arch    string
		Version string
	}

	type appInfo struct {
		Name    string
		Version string
	}
	osArr := []osInfo{
		{
			Name:    "linux",
			Arch:    "amd64",
			Version: "ubuntu 22.04",
		},
		{
			Name:    "linux",
			Arch:    "amd64",
			Version: "ubuntu 18.06",
		},
		{
			Name:    "windows",
			Arch:    "amd64",
			Version: "10",
		},
		{
			Name:    "darwin",
			Arch:    "amd64",
			Version: "10.0.3",
		},
		{
			Name:    "darwin",
			Arch:    "amd64",
			Version: "10.0.1",
		},
	}
	appsArr := []appInfo{
		{
			Name:    "financial",
			Version: "v0.0.1",
		},
		{
			Name:    "financial",
			Version: "v1.0.1",
		},
		{
			Name:    "sports",
			Version: "v2.1.1",
		},
		{
			Name:    "sports",
			Version: "v1.1.1",
		},
		{
			Name:    "educational",
			Version: "v0.1.1",
		},
		{
			Name:    "educational",
			Version: "v3.1.1",
		},
	}
	var result []entities.CoreDump
	for i := 0; i < quantity; i++ {
		coreDump := entities.NewCoreDump()
		coreDump.ID = fmt.Sprint(i)
		osInfo := entities.NewOSInfo()
		osInfo.SetName(osArr[i].Name)
		osInfo.SetArchitecture(osArr[i].Arch)
		osInfo.SetVersion(osArr[i].Version)
		coreDump.SetOSInfo(osInfo)

		appInfo := entities.NewAppInfo()
		appInfo.SetName(appsArr[i].Name)
		appInfo.SetProgrammingLanguage(entities.ProgrammingLanguage(1))
		appInfo.SetVersion(appsArr[i].Version)
		coreDump.SetAppInfo(appInfo)

		coreDump.SetStatus(1)
		coreDump.SetData(time.Now().Format("2006-01-02"))

		coreDump.SetTimestamp(time.Unix(1663511325, 0))

		result = append(result, *coreDump)
	}

	return result
}
