package note

import (
	"context"
	"fmt"

	"github.com/scipie28/note-service-api/internal/app/model"
)

func (n *note) DescribeNote(ctx context.Context, id int64) (*model.Note, error) {
	fmt.Printf("description of note with id: %v", id)

	return &model.Note{
		Id:          33,
		UserId:      33,
		ClassroomId: 33,
		DocumentId:  33,
	}, nil
}
