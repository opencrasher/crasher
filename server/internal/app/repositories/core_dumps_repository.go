package repositories

import (
	"server/internal/app/entities"
)

type CoreDumpsRepository interface {
	GetCoreDumps(parameters ...interface{}) ([]entities.CoreDump, error)
	GetCoreDumpByID(id string) (entities.CoreDump, error)
	AddCoreDump(coreDump entities.CoreDump) error
	UpdateCoreDump(parameters ...interface{}) error
	DeleteCoreDump(id string) error
	DeleteAllCoreDumps() error
}
