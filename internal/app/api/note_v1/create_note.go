package note_v1

import (
	"context"

	"github.com/scipie28/note-service-api/internal/app/model"
	pb "github.com/scipie28/note-service-api/pkg/note_v1"
)

func (n *Note) CreateNoteV1(ctx context.Context, req *pb.CreateNoteV1Request) (*pb.CreateNoteV1Response, error) {
	id, err := n.noteService.CreateNote(ctx, &model.Note{
		UserId:      req.GetUserId(),
		ClassroomId: req.GetClassroomId(),
		DocumentId:  req.GetDocumentId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateNoteV1Response{NoteId: id}, nil
}
