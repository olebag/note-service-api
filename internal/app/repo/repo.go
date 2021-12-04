package repo

//go:generate mockgen -destination=mocks/mock_repo.go -package=mocks . Repo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/scipie28/note-service-api/internal/app/model"
)

const tableName = "notes"

type Repo interface {
	AddNote(ctx context.Context, note *model.Note) (int64, error)
	MultiAddNotes(notes []model.Note) (int64, error)
	UpdateNote(id int64, note model.Note) error
	RemoveNote(id int64) error
	DescribeNote(id int64) (model.Note, error)
	ListNotes() ([]*model.Note, error)
}

type repo struct {
	db sqlx.DB
}

func NewRepo(db sqlx.DB) Repo {
	return &repo{
		db: db,
	}
}

func (r *repo) AddNote(ctx context.Context, note *model.Note) (int64, error) {
	q := sq.Insert(tableName).
		Columns("user_id", "classroom_id", "document_id").
		Values(note.UserId, note.ClassroomId, note.DocumentId).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING \"id\"")

	s, _, _ := q.ToSql()
	fmt.Println(s)

	err := q.QueryRowContext(ctx).Scan(&note.Id)
	if err != nil {
		return 0, err
	}

	return note.Id, nil
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

func (r *repo) ListNotes() ([]*model.Note, error) {
	data := []*model.Note{
		{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
		{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
		{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
		{Id: 5, UserId: 5, ClassroomId: 23, DocumentId: 6},
		{Id: 555, UserId: 6, ClassroomId: 24, DocumentId: 7},
	}

	return data, nil
}
