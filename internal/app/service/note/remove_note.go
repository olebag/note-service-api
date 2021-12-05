package note

import (
	"context"
)

func (n *note) RemoveNote(ctx context.Context, id int64) error {
	return n.noteRepo.RemoveNote(ctx, id)
}
