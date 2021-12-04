package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/scipie28/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
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

	resAdd, err := client.AddNoteV1(ctx, &pb.AddNoteV1Request{
		UserId:      2435,
		ClassroomId: 223,
		DocumentId:  3123,
	})
	if err != nil {
		log.Fatalf("failed to adding: %s", err.Error())
	}

	fmt.Printf("added note with id: %s\n\n", resAdd)
	//
	//desc, err := client.DescribeNoteV1(ctx, &pb.DescribeNoteV1Request{Id: 333})
	//if err != nil {
	//	log.Fatalf("failed to describing: %s", err.Error())
	//}
	//
	//fmt.Printf("description of note with id: %v\n\n", desc)
	//
	//upd, err := client.UpdateNoteV1(ctx, &pb.UpdateNoteV1Request{
	//	Id:          999,
	//	UserId:      9999,
	//	ClassroomId: 99999,
	//	DocumentId:  999999,
	//})
	//if err != nil {
	//	log.Fatalf("failid to updating: %s", err.Error())
	//}
	//
	//fmt.Printf("note with id %v updated\n\n", upd)
	//
	//resRem, err := client.RemoveNoteV1(ctx, &pb.RemoveNoteV1Request{
	//	Id: 123,
	//})
	//if err != nil {
	//	log.Fatalf("failid to removing: %s", err.Error())
	//}
	//
	//fmt.Printf("note with id %v deleted\n\n", resRem)
}
