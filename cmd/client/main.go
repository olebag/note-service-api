package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/scipie28/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const address = "localhost:50051"

func main() {
	ctx := context.Background()

	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer func() {
		err = con.Close()
		if err != nil {
			log.Fatalf("failed to closing connect")
		}
	}()

	client := pb.NewNoteV1Client(con)

	add(ctx, client)
	multiAddNotes(ctx, client)
	describeNote(ctx, client)
	listNotes(ctx, client)
	removeNote(ctx, client)
	updateNote(ctx, client)
}

func add(ctx context.Context, client pb.NoteV1Client) {
	res, err := client.AddNoteV1(ctx, &pb.AddNoteV1Request{
		UserId:      21,
		ClassroomId: 20,
		DocumentId:  31,
	})
	if err != nil {
		log.Fatalf("failed to adding: %s", err.Error())
	}

	fmt.Printf("added note with id: %s\n\n", res)
}

func multiAddNotes(ctx context.Context, client pb.NoteV1Client) {
	var notes = []*pb.Notes{
		{Id: 11, UserId: 1, ClassroomId: 23, DocumentId: 66000},
		{Id: 21, UserId: 245, ClassroomId: 24, DocumentId: 3677000},
		{Id: 31, UserId: 33, ClassroomId: 23, DocumentId: 86000},
		{Id: 41, UserId: 74, ClassroomId: 234, DocumentId: 7000},
		{Id: 51, UserId: 50, ClassroomId: 23, DocumentId: 26000},
		{Id: 5155, UserId: 56, ClassroomId: 244, DocumentId: 367000},
	}

	res, err := client.MultiAddNotesV1(ctx, &pb.MultiAddNotesV1Request{
		Notes: notes,
	})
	if err != nil {
		log.Fatalf("failed to milti adding: %s", err.Error())
	}

	fmt.Printf("adding %v notes", res)
}

func describeNote(ctx context.Context, client pb.NoteV1Client) {
	res, err := client.DescribeNoteV1(ctx, &pb.DescribeNoteV1Request{Id: 1})
	if err != nil {
		log.Fatalf("failed to describtion: %s", err.Error())
	}

	fmt.Printf("id: %v, user_id: %v, classroom_id: %v, document_id: %v", res.Id, res.UserId, res.ClassroomId, res.DocumentId)
}

func listNotes(ctx context.Context, client pb.NoteV1Client) {
	res, err := client.ListNotesV1(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("failed to listing notes: %s", err.Error())
	}

	for _, note := range res.Notes {
		fmt.Printf("id: %v, user_id: %v, classroom_id: %v, document_id: %v\n",
			note.Id, note.UserId, note.ClassroomId, note.DocumentId)
	}
}

func removeNote(ctx context.Context, client pb.NoteV1Client) {
	_, err := client.RemoveNoteV1(ctx, &pb.RemoveNoteV1Request{Id: 15})
	if err != nil {
		log.Fatalf("failed to removing note: %s", err.Error())
	}
}

func updateNote(ctx context.Context, client pb.NoteV1Client) {
	_, err := client.UpdateNoteV1(ctx, &pb.UpdateNoteV1Request{
		Id:          5,
		UserId:      999,
		ClassroomId: 999,
		DocumentId:  999,
	})

	if err != nil {
		log.Fatalf("failed to updating note: %s", err.Error())
	}
}
