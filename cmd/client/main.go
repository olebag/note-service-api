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

	resAdd, err := client.AddNoteV1(ctx, &pb.AddNoteV1Request{
		UserId:      1,
		ClassroomId: 2,
		DocumentId:  3,
	})
	if err != nil {
		log.Fatalf("failid to adding: %s", err.Error())
	}

	fmt.Printf("Method Add. Id response: %s\n\n", resAdd)

	resRem, err := client.RemoveNoteV1(ctx, &pb.RemoveNoteV1Request{Id: 123})
	if err != nil {
		log.Fatalf("failid to removing: %s", err.Error())
	}

	fmt.Printf("Method Remove. Respone: %v\n\n", resRem)

	resDescr, err := client.DescribeNoteV1(ctx, &pb.DescribeNoteV1Request{Id: 333})
	if err != nil {
		log.Fatalf("failid to describing: %s", err.Error())
	}
	fmt.Printf("Method Describe. Respone: %v\n\n", resDescr)

	resUpd, err := client.UpdateNoteV1(ctx, &pb.UpdateNoteV1Request{
		Id:          999,
		UserId:      9999,
		ClassroomId: 99999,
		DocumentId:  999999,
	})
	if err != nil {
		log.Fatalf("failid to updating: %s", err.Error())
	}

	fmt.Printf("Method Update. Response: %v\n\n", resUpd)

}
