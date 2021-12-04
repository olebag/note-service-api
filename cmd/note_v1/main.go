package main

import (
	"fmt"
	"log"
	"net"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/scipie28/note-service-api/internal/app/api/note_v1"
	"github.com/scipie28/note-service-api/internal/app/repo"
	"github.com/scipie28/note-service-api/internal/app/service/note"
	pb "github.com/scipie28/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	dbInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		"localhost", 54321, "shem", "12345678", "note-service-api", "disable")

	db, err := sqlx.Open("pgx", dbInfo)
	if err != nil {
		log.Fatalf("failed to opening connection to db: %s", err.Error())
	}
	defer db.Close()

	repo := repo.NewRepo(*db)

	s := grpc.NewServer()
	pb.RegisterNoteV1Server(s, &note_v1.Note{NoteService: note.NewNote(repo)})

	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to server: %s", err.Error())
	}
}
