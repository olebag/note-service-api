package note_v1

import (
	"context"
	pb "github.com/scipie28/note-service-api/pkg/note_v1"
)

func (n *Note) DescribeV1(ctx context.Context, req *pb.DescribeNoteV1Request) (res *pb.DescribeNoteV1Response, err error) {
	note, err := n.NoteService.Describe(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.DescribeNoteV1Response{
		UserId:      note.UserId,
		ClassroomId: note.ClassroomId,
		DocumentId:  note.DocumentId,
	}, nil
}
