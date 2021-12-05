package note

import (
	"context"

	"github.com/scipie28/note-service-api/internal/app/model"
)

func (n *note) ListNotes(ctx context.Context) ([]*model.Note, error) {
	return n.noteRepo.ListNotes(ctx)
}
