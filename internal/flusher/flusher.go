package flusher

import (
	"log"

	"github.com/scipie28/note-service-api/internal/app/api"
	"github.com/scipie28/note-service-api/internal/repo"
	mocksRepo "github.com/scipie28/note-service-api/internal/repo/mocks"
	"github.com/scipie28/note-service-api/internal/utills"
)

type Flusher interface {
	Flush(note []api.Note, batchSize uint64) ([]api.Note, error)
}

type flusher struct {
	repo repo.Repo
}

func NewFlusher(repo *mocksRepo.MockRepo) Flusher {
	return &flusher{repo}
}

func (f *flusher) Flush(notes []api.Note, batchSize uint64) ([]api.Note, error) {
	batches, err := utills.SplitSlice(notes, batchSize)
	if err != nil {
		log.Printf("failed to spliting slice: %s", err.Error())
		return nil, err
	}

	for i, batch := range batches {
		num, err := f.repo.MultiAdd(batch)
		if err != nil {
			log.Printf("failed to add slice: %s", err.Error())

			var save = make([]api.Note, 0)
			for _, v := range batches[i:] {
				save = append(save, v...)
			}

			return save, err
		}

		log.Printf("%d notes added", num)
	}

	return nil, nil
}
