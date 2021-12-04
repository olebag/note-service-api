package note

import (
	"context"
	"github.com/scipie28/note-service-api/internal/app/model"
)

func (n *note) ListNotes(ctx context.Context) ([]*model.Note, error) {
	notes := []*model.Note{
		{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
		{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
		{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
		{Id: 5, UserId: 5, ClassroomId: 23, DocumentId: 6},
		{Id: 555, UserId: 6, ClassroomId: 24, DocumentId: 7},
	}

	return notes, nil
}
