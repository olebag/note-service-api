package note_v1

import (
	"context"
	"github.com/scipie28/note-service-api/internal/app/model"
)
import pb "github.com/scipie28/note-service-api/pkg/note_v1"

func (n *Note) UpdateV1(ctx context.Context, req *pb.UpdateNoteV1Request) error {
	err := n.NoteService.Update(ctx, &model.Note{
		Id:          req.GetId(),
		UserId:      req.GetUserId(),
		ClassroomId: req.GetClassroomId(),
		DocumentId:  req.DocumentId,
	})
	if err != nil {
		return err
	}

	return nil
}
