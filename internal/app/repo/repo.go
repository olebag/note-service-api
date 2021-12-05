package repo

//go:generate mockgen -destination=mocks/mock_repo.go -package=mocks . Repo

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/scipie28/note-service-api/internal/app/model"
)

const (
	limit     = 10
	tableName = "notes"
)

type Repo interface {
	AddNote(ctx context.Context, note *model.Note) (int64, error)
	MultiAddNotes(ctx context.Context, notes []*model.Note) (int64, error)
	UpdateNote(ctx context.Context, id int64, note *model.Note) error
	RemoveNote(ctx context.Context, id int64) error
	DescribeNote(ctx context.Context, id int64) (*model.Note, error)
	ListNotes(ctx context.Context) ([]*model.Note, error)
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

func (r *repo) MultiAddNotes(ctx context.Context, notes []*model.Note) (int64, error) {
	var num int64
	q := sq.Insert(tableName).
		Columns("user_id", "classroom_id", "document_id").
		RunWith(r.db).PlaceholderFormat(sq.Dollar)

	for _, note := range notes {
		q = q.Values(note.UserId, note.ClassroomId, note.DocumentId)
		num++
	}

	_, err := q.ExecContext(ctx)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func (r *repo) UpdateNote(ctx context.Context, id int64, note *model.Note) error {
	q := sq.Update(tableName).SetMap(map[string]interface{}{
		"user_id":      note.UserId,
		"classroom_id": note.ClassroomId,
		"document_id":  note.DocumentId,
	}).
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	_, err := q.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) RemoveNote(ctx context.Context, id int64) error {
	q := sq.Delete(tableName).
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	_, err := q.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) DescribeNote(ctx context.Context, id int64) (*model.Note, error) {
	var note model.Note
	q := sq.Select("user_id", "classroom_id", "document_id").
		From(tableName).
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	err := q.QueryRowContext(ctx).
		Scan(&note.UserId, &note.ClassroomId, &note.DocumentId)

	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *repo) ListNotes(ctx context.Context) ([]*model.Note, error) {
	var res []*model.Note
	q := sq.Select("id", "user_id", "classroom_id", "document_id").
		From(tableName).
		RunWith(r.db).
		Limit(limit).
		Offset(1).
		PlaceholderFormat(sq.Dollar)

	rows, err := q.QueryContext(ctx)
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Fatalf("failed to closing rows: %s", err.Error())
		}
	}(rows)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id, userId, classroomId, documentId int64

		if err = rows.Scan(&id, &userId, &classroomId, &documentId); err != nil {
			return nil, err
		}

		res = append(res, &model.Note{
			Id:          id,
			UserId:      userId,
			ClassroomId: classroomId,
			DocumentId:  documentId,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
