package note

import (
	"context"

	"github.com/scipie28/note-service-api/internal/app/model"
)

type INote interface {
	AddNote(ctx context.Context, note *model.Note) (int64, error)
	RemoveNote(ctx context.Context, id int64) error
	UpdateNote(ctx context.Context, note *model.Note) error
	DescribeNote(ctx context.Context, id int64) (*model.Note, error)
}

type note struct {
}

func NewNote() INote {
	return &note{}
}
