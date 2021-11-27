package note

import (
	"context"
	"github.com/scipie28/note-service-api/internal/app/model"
)

type INote interface {
	CreateNote(ctx context.Context, note *model.Note) (int64, error)
}

type note struct {
}

func NewNote() INote {
	return &note{}
}