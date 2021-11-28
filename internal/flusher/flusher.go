package flusher

import (
	"log"

	"github.com/scipie28/note-service-api/internal/app/model"
	"github.com/scipie28/note-service-api/internal/repo"
	"github.com/scipie28/note-service-api/internal/utills"
)

type Flusher interface {
	Flush(note []model.Note, batchSize int64) ([]model.Note, error)
}

type flusher struct {
	repo repo.Repo
}

func NewFlusher(repo repo.Repo) Flusher {
	return &flusher{repo}
}

func (f *flusher) Flush(notes []model.Note, batchSize int64) ([]model.Note, error) {
	batches, err := utills.SplitSlice(notes, batchSize)
	if err != nil {
		log.Printf("failed to spliting slice: %s", err.Error())
		return nil, err
	}

	for i, batch := range batches {
		num, errAdd := f.repo.MultiAddNotes(batch)
		if errAdd != nil {
			log.Printf("failed to adding slice: %s", errAdd.Error())

			var save = make([]model.Note, 0)
			for _, v := range batches[i:] {
				save = append(save, v...)
			}

			return save, errAdd
		}

		log.Printf("%d notes added", num)
	}

	return nil, nil
}
