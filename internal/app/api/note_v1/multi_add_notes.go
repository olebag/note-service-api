package note_v1

import (
	"context"

	"github.com/scipie28/note-service-api/internal/app/model"
	pb "github.com/scipie28/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) MultiAddNotesV1(ctx context.Context, req *pb.MultiAddNotesV1Request) (*emptypb.Empty, error) {
	var notes []*model.Note

	for _, note := range req.GetNotes() {
		notes = append(notes, &model.Note{
			Id:          note.GetId(),
			UserId:      note.GetUserId(),
			ClassroomId: note.GetClassroomId(),
			DocumentId:  note.GetDocumentId(),
		})
	}

	err := n.NoteService.MultiAddNotes(ctx, notes)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
