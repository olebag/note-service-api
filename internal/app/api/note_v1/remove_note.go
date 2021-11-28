package note_v1

import (
	"context"
	pb "github.com/scipie28/note-service-api/pkg/note_v1"
)

func (n *Note) RemoveV1(ctx context.Context, req *pb.RemoveNoteV1Request) error {
	err := n.NoteService.Remove(ctx, req.GetId())
	if err != nil {
		return err
	}

	return nil
}
