package note

import (
	"context"
	"fmt"
	"github.com/scipie28/note-service-api/internal/app/model"
)

func (n *note) Describe(ctx context.Context, id int64) (*model.Note, error) {
	fmt.Printf("Printing note with id %v", id)

	return &model.Note{
		Id:          33,
		UserId:      33,
		ClassroomId: 33,
		DocumentId:  33,
	}, nil
}
