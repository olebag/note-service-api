package note

import (
	"context"

	"github.com/scipie28/note-service-api/internal/app/model"
)

func (n *note) UpdateNote(ctx context.Context, note *model.Note) error {
	return n.noteRepo.UpdateNote(ctx, note.Id, note)
}
