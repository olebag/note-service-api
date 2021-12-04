package note_v1

import (
	"context"
	pb "github.com/scipie28/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) ListNotesV1(ctx context.Context, _ *emptypb.Empty) (*pb.ListNotesV1Response, error) {
	var res []*pb.Notes

	notes, err := n.NoteService.ListNotes(ctx)
	if err != nil {
		return nil, err
	}

	for _, note := range notes {
		res = append(res, &pb.Notes{
			Id:          note.Id,
			UserId:      note.UserId,
			ClassroomId: note.ClassroomId,
			DocumentId:  note.DocumentId,
		})
	}

	return &pb.ListNotesV1Response{Notes: res}, nil
}
