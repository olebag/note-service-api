package repo

//go:generate mockgen -destination=mocks/mock_repo.go -package=mocks . Repo

import (
	"github.com/scipie28/note-service-api/internal/app/model"
)

type Repo interface {
	Add(note model.Note) error
	MultiAdd(notes []model.Note) (int64, error)
	Update(id int64, note model.Note) error
	Remove(id int64) error
	Describe(id int64) (model.Note, error)
}
