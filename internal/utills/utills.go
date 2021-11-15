package utills

import (
	"errors"
	"fmt"

	"github.com/scipie28/note-service-api/internal/app/api"
)

var filter = []string{"a", "b", "c", "d"}

func SwapKeyAndValue(data map[int64]string) (map[string]int64, error) {
	res := make(map[string]int64)

	for key, value := range data {
		if _, found := res[value]; found {
			fmt.Println("The key exists", value)
			continue
		}

		res[value] = key
	}

	return res, nil
}

func FilterSlice(data []string) ([]string, error) {
	dataMap := make(map[string]struct{})
	var res []string

	for _, val := range data {
		dataMap[val] = struct{}{}
	}

	for _, val := range filter {
		if _, found := dataMap[val]; found {
			res = append(res, val)
		}
	}

	return res, nil
}

func ConvertSliceToMap(users []api.Note) (map[uint64]api.Note, error) {
	res := make(map[uint64]api.Note)
	for _, v := range users {
		res[v.Id] = v
	}

	return res, nil
}

func SplitSlice(notes []api.Note, batchSize uint64) ([][]api.Note, error) {
	if batchSize <= 0 || notes == nil {
		return nil, errors.New("error input values")
	}

	if uint64(len(notes)) <= batchSize {
		return [][]api.Note{notes}, nil
	}

	numBatches := uint64(len(notes)) / batchSize
	if uint64(len(notes))%batchSize != 0 {
		numBatches++
	}

	var end uint64

	res := make([][]api.Note, 0, numBatches)

	for begin := uint64(0); begin < uint64(len(notes)); {
		end += batchSize
		if end > uint64(len(notes)) {
			end = uint64(len(notes))
		}

		res = append(res, notes[begin:end])
		begin += batchSize
	}

	return res, nil
}
