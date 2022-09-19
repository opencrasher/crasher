package services

import (
	"server/internal/app/entities"
	"server/internal/app/repositories"

	"go.uber.org/zap"
)

type CoreDumpsService interface {
	GetCoreDumps(parameters ...interface{}) ([]entities.CoreDump, error)
	GetCoreDumpByID(id string) (entities.CoreDump, error)
	AddCoreDump(coreDump entities.CoreDump) error
	UpdateCoreDump(parameters ...interface{}) error
	DeleteCoreDump(id string) error
	DeleteAllCoreDumps() error
}

type CoreDumpServiceImpl struct {
	repository repositories.CoreDumpsRepository
	logger     *zap.Logger
}

func NewCoreDumpsService(r repositories.CoreDumpsRepository, l *zap.Logger) CoreDumpsService {
	return &CoreDumpServiceImpl{
		repository: r,
		logger:     l,
	}
}

func (s *CoreDumpServiceImpl) GetCoreDumps(parameters ...interface{}) ([]entities.CoreDump, error) {
	return nil, nil
}

func (s *CoreDumpServiceImpl) GetCoreDumpByID(id string) (entities.CoreDump, error) {
	return entities.CoreDump{}, nil
}

func (s *CoreDumpServiceImpl) AddCoreDump(coreDump entities.CoreDump) error {
	return nil
}

func (r *CoreDumpServiceImpl) UpdateCoreDump(parameters ...interface{}) error {
	return nil
}

func (s *CoreDumpServiceImpl) DeleteCoreDump(id string) error {
	return nil
}

func (s *CoreDumpServiceImpl) DeleteAllCoreDumps() error {
	return nil
}
