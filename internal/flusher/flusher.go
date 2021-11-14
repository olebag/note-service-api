package flusher

import (
	"errors"
	"fmt"
	"github.com/scipie28/note-service-api/internal/app/api"
	"github.com/scipie28/note-service-api/internal/repo"
	"github.com/scipie28/note-service-api/internal/utills"
)

type Flusher interface {
	Flush(note []api.Note, butch uint32) ([][]api.Note, error)
}

type flusher struct {
	addData repo.Repo
}

func New(data repo.Repo) Flusher {
	return &flusher{data}
}

func (f *flusher) Flush(data []api.Note, butchSize uint32) ([][]api.Note, error) {
	chanks, _ := utills.SplitSlice(data, butchSize)
	for i := 0; i < len(chanks); i++ {
		err := f.addData.Add(chanks[i])

		if err != nil {
			fmt.Println(err)
			save := chanks[i:]
			return save, errors.New("transmission was cut off")
		}
	}

	return nil, nil
}
