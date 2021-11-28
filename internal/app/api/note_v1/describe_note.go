package note_v1

import (
	"context"
	pb "github.com/scipie28/note-service-api/pkg/note_v1"
)

func (n *Note) DescribeNoteV1(ctx context.Context, req *pb.DescribeNoteV1Request) (*pb.DescribeNoteV1Response, error) {
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
