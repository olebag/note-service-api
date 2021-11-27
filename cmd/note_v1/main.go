package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"github.com/scipie28/note-service-api/internal/app/api/note_v1"
	pb "github.com/scipie28/note-service-api/pkg/note_v1"
)

const port = "50051"

func main(){
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterNoteV1Server(s, &note_v1.Note{})
}

