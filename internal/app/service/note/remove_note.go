package note

import (
	"context"
	"fmt"
)

func (n *note) Remove(ctx context.Context, id int64) error {
	fmt.Printf("Note with id %v removing", id)
	return nil
}
