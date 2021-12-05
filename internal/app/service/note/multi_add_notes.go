package note

import (
	"context"

	"github.com/scipie28/note-service-api/internal/app/model"
)

func (n *note) MultiAddNotes(ctx context.Context, notes []*model.Note) (int64, error) {
	return n.noteRepo.MultiAddNotes(ctx, notes)
}
