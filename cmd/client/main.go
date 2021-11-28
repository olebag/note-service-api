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

	resAdd, _ := client.AddNoteV1(ctx, &pb.AddNoteV1Request{
		UserId:      1,
		ClassroomId: 2,
		DocumentId:  3,
	})
	fmt.Printf("Method Add. Id response: %s\n\n", resAdd)

	resRem, _ := client.RemoveNoteV1(ctx, &pb.RemoveNoteV1Request{Id: 123})
	fmt.Printf("Method Remove. Respone: %v\n\n", resRem)

	resDescr, _ := client.DescribeNoteV1(ctx, &pb.DescribeNoteV1Request{Id: 333})
	fmt.Printf("Method Describe. Respone: %v\n\n", resDescr)

	resUpd, _ := client.UpdateNoteV1(ctx, &pb.UpdateNoteV1Request{
		Id:          999,
		UserId:      9999,
		ClassroomId: 99999,
		DocumentId:  999999,
	})
	fmt.Printf("Method Update. Response: %v\n\n", resUpd)

}
