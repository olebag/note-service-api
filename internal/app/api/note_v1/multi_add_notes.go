package note_v1

import (
	"context"

	"github.com/scipie28/note-service-api/internal/app/model"
	pb "github.com/scipie28/note-service-api/pkg/note_v1"
)

func (n *Note) MultiAddNotesV1(ctx context.Context, req *pb.MultiAddNotesV1Request) (*pb.MultiAddNotesV1Response, error) {
	var notes []*model.Note

	for _, note := range req.GetNotes() {
		notes = append(notes, &model.Note{
			Id:          note.GetId(),
			UserId:      note.GetUserId(),
			ClassroomId: note.GetClassroomId(),
			DocumentId:  note.GetDocumentId(),
		})
	}

	res, err := n.NoteService.MultiAddNotes(ctx, notes)
	if err != nil {
		return nil, err
	}

	return &pb.MultiAddNotesV1Response{Id: res}, nil
}
