package repo

import (
	"errors"
	"fmt"
	"github.com/scipie28/note-service-api/internal/app/api"
)

type Repo interface {
	Add([]api.Note) error
	Delete(string2 string)
}

type repo struct {
	num int32
}

func New(num int32) Repo {
	return &repo{num: num}
}

func (r *repo) Delete(string) {
	fmt.Printf("")
}

func (r *repo) Add(data []api.Note) error {
	r.num++
	fmt.Println("write - ", data)
	if r.num > 2 {
		return errors.New("FatalError")
	}
	return nil
}

//mockgen -destination=mocks/mock_repo.go -package=mocks github.com/scipie28/note-service-api/cmd/note-service-api/iternal/repo repo
