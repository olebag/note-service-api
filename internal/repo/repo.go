package repo

//go:generate mockgen -destination=mocks/mock_repo.go -package=mocks . Repo

import (
	"fmt"
	"github.com/scipie28/note-service-api/internal/app/model"
)

type Repo interface {
	Create(id, UserId, ClassroomId, DocumentId  int64) int64
	Add(note model.Note) error
	MultiAdd(notes []model.Note) (int64, error)
	Update(id int64, note model.Note) error
	Remove(id int64) error
	Describe(id int64) (model.Note, error)
}

type repo struct {
}

func NewRepo() Repo {
	return &repo{}
}

func (r *repo) Create(id, UserId, ClassroomId, DocumentId int64) int64{
	note := model.Note{Id: id, UserId: UserId, ClassroomId: ClassroomId, DocumentId: DocumentId}
	fmt.Printf("note was created:")
	note.String()
	return note.Id
}


func (r *repo) Add(note model.Note) error {
	panic("implement me")
}

func (r *repo) MultiAdd(notes []model.Note) (int64, error) {
	fmt.Println(notes)
	return int64(len(notes)), nil
}

func (r *repo) Update(id int64, note model.Note) error {
	panic("implement me")
}

func (r *repo) Remove(id int64) error {
	panic("implement me")
}

func (r *repo) Describe(id int64) (model.Note, error) {
	panic("implement me")
}
