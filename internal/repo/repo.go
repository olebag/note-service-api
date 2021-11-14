package repo

//go:generate mockgen -destination=mocks/mock_repo.go -package=mocks . Repo

import (
	"github.com/scipie28/note-service-api/internal/app/api"
)

type Repo interface {
	Add(api.Note) error
	MultiAdd([]api.Note) (int32, error)      //окл-во записей записанных и ошибко
	Update(int32, api.Note) error            // индекс, ноту/ ерор
	Remove(int32) error                      // ид/ еррор
	Describe(int322 int32) (api.Note, error) //  приним ид/ саму ноду-которую записал и  еррор
}
