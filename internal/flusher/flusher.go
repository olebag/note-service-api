package flusher

import (
	"log"

	"github.com/scipie28/note-service-api/internal/app/api"
	"github.com/scipie28/note-service-api/internal/repo"
	"github.com/scipie28/note-service-api/internal/utills"
)

type Flusher interface {
	Flush(note []api.Note, butchSize uint32) ([]api.Note, error)
}

type flusher struct {
	repo repo.Repo
}

func New(repo repo.Repo) Flusher {
	return &flusher{repo}
}

func (f *flusher) Flush(notes []api.Note, batchSize uint32) ([]api.Note, error) {
	batchs, err := utills.SplitSlice(notes, batchSize)
	if err != nil {
		return nil, err
	}

	for i, batch := range batchs {
		num, err := f.repo.MultiAdd(batch)
		if err != nil {
			save, err2 := utills.TwoToOneDimensionalSlice(batchs[i:])

			if err2 != nil {
				return nil, err2
			}

			return save, err
		}
		log.Printf("%d notes added", num)
	}

	return nil, nil
}
