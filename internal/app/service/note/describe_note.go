package note

import (
	"context"

	"github.com/scipie28/note-service-api/internal/app/model"
)

func (n *note) DescribeNote(ctx context.Context, id int64) (*model.Note, error) {
	return n.noteRepo.DescribeNote(ctx, id)
}
