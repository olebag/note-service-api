package note_v1

import (
	"context"

	"github.com/scipie28/note-service-api/internal/app/model"
	pb "github.com/scipie28/note-service-api/pkg/note_v1"
)

func (n *Note) AddNoteV1(ctx context.Context, req *pb.AddNoteV1Request) (*pb.AddNoteV1Response, error) {
	id, err := n.NoteService.AddNote(ctx, &model.Note{
		UserId:      req.GetUserId(),
		ClassroomId: req.GetClassroomId(),
		DocumentId:  req.GetDocumentId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddNoteV1Response{NoteId: id}, nil
}
