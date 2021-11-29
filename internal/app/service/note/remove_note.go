package note

import (
	"context"
	"fmt"
)

func (n *note) RemoveNote(ctx context.Context, id int64) error {
	fmt.Printf("note with id %v deleted", id)

	return nil
}
