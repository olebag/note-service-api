package note_v1

import (
	"context"
	pb "github.com/scipie28/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) RemoveNoteV1(ctx context.Context, req *pb.RemoveNoteV1Request) (*emptypb.Empty, error) {
	err := n.NoteService.Remove(ctx, req.GetId())
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
