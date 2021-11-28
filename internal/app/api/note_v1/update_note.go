package note_v1

import (
	"context"

	"github.com/scipie28/note-service-api/internal/app/model"
	"google.golang.org/protobuf/types/known/emptypb"
)
import pb "github.com/scipie28/note-service-api/pkg/note_v1"

func (n *Note) UpdateNoteV1(ctx context.Context, req *pb.UpdateNoteV1Request) (*emptypb.Empty, error) {
	err := n.NoteService.UpdateNote(ctx, &model.Note{
		Id:          req.GetId(),
		UserId:      req.GetUserId(),
		ClassroomId: req.GetClassroomId(),
		DocumentId:  req.GetDocumentId(),
	})
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
