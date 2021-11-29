package repo

//go:generate mockgen -destination=mocks/mock_repo.go -package=mocks . Repo

import (
	"fmt"

	"github.com/scipie28/note-service-api/internal/app/model"
)

type Repo interface {
	AddNote(note model.Note) error
	MultiAddNotes(notes []model.Note) (int64, error)
	UpdateNote(id int64, note model.Note) error
	RemoveNote(id int64) error
	DescribeNote(id int64) (model.Note, error)
}

type repo struct {
}

func NewRepo() Repo {
	return &repo{}
}

func (r *repo) AddNote(note model.Note) error {
	fmt.Println(note)

	return nil
}

func (r *repo) MultiAddNotes(notes []model.Note) (int64, error) {
	fmt.Println(notes)

	return int64(len(notes)), nil
}

func (r *repo) UpdateNote(id int64, note model.Note) error {
	fmt.Println(note, id)

	return nil
}

func (r *repo) RemoveNote(id int64) error {
	fmt.Println(id)

	return nil
}

func (r *repo) DescribeNote(id int64) (model.Note, error) {
	fmt.Println(id)

	return model.Note{
		Id:          1,
		UserId:      1,
		ClassroomId: 0,
		DocumentId:  0,
	}, nil
}
