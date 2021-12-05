package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/scipie28/note-service-api/internal/app/api/note_v1"
	"github.com/scipie28/note-service-api/internal/app/repo"
	"github.com/scipie28/note-service-api/internal/app/service/note"
	pb "github.com/scipie28/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
)

const portGrpc = ":50051"
const portHttp = ":8090"

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		log.Fatal(startGRPC())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		log.Fatal(startREST())
		wg.Done()
	}()

	wg.Wait()
}

func startGRPC() error {
	dbInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		"localhost", 54321, "shem", "12345678", "note-service-api", "disable")

	db, err := sqlx.Open("pgx", dbInfo)
	if err != nil {
		log.Fatalf("failed to opening connection to db: %s", err.Error())
	}
	defer db.Close()

	repo := repo.NewRepo(*db)

	lis, err := net.Listen("tcp", portGrpc)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNoteV1Server(grpcServer, &note_v1.Note{NoteService: note.NewNote(repo)})
	grpcServer.Serve(lis)

	return nil
}

func startREST() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterNoteV1HandlerFromEndpoint(ctx, mux, portGrpc, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(portHttp, mux)
}
