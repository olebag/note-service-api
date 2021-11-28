package note

import (
	"context"
	"fmt"
	"github.com/scipie28/note-service-api/internal/app/model"
)

func (n *note) Update(ctx context.Context, note *model.Note) error {
	fmt.Printf("Note with id %v updateing: %v", note.Id, note)
	return nil
}
