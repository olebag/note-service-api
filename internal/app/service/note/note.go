package note

import (
	"context"

	"github.com/scipie28/note-service-api/internal/app/model"
	"github.com/scipie28/note-service-api/internal/app/repo"
)

type INote interface {
	AddNote(ctx context.Context, note *model.Note) (int64, error)
	RemoveNote(ctx context.Context, id int64) error
	UpdateNote(ctx context.Context, note *model.Note) error
	DescribeNote(ctx context.Context, id int64) (*model.Note, error)
	MultiAddNotes(ctx context.Context, notes []*model.Note) (int64, error)
	ListNotes(ctx context.Context) ([]*model.Note, error)
}

type note struct {
	noteRepo repo.Repo
}

func NewNote(noteRepo repo.Repo) INote {
	return &note{
		noteRepo: noteRepo,
	}
}
