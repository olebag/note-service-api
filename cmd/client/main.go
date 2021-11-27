package main

import (
	"context"
	"fmt"
	pb "github.com/scipie28/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
	"log"
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

	res, err := client.CreateNoteV1(ctx, &pb.CreateNoteV1Request{
		UserId:      1,
		ClassroomId: 2,
		DocumentId:  3,
	})

	fmt.Printf("Id response: %s", res)
}
