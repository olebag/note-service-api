package note

import (
	"context"
	"fmt"

	"github.com/scipie28/note-service-api/internal/app/model"
)

func (n *note) UpdateNote(ctx context.Context, note *model.Note) error {
	fmt.Printf("note %v with id %v updated ", note, note.Id)

	return nil
}
