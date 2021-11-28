package note

import (
	"context"
	"github.com/scipie28/note-service-api/internal/app/model"
)

func (n *note) AddNote(ctx context.Context, note *model.Note) (int64, error) {
	note.String()

	return 1, nil
}
