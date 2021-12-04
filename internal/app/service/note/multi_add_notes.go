package note

import (
	"context"

	"github.com/scipie28/note-service-api/internal/app/model"
)

func (n *note) MultiAddNotes(ctx context.Context, notes []*model.Note) error {
	for _, nt := range notes {
		nt.String()
	}

	return nil
}
