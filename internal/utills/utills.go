package utills

import (
	"errors"
	"fmt"

	"github.com/scipie28/note-service-api/internal/app/model"
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

func ConvertSliceToMap(users []model.Note) (map[int64]model.Note, error) {
	res := make(map[int64]model.Note)
	for _, v := range users {
		res[v.Id] = v
	}

	return res, nil
}

func SplitSlice(notes []model.Note, batchSize int64) ([][]model.Note, error) {
	if batchSize <= 0 || notes == nil || len(notes) == 0 {
		return nil, errors.New("invalid input values")
	}

	if int64(len(notes)) <= batchSize {
		return [][]model.Note{notes}, nil
	}

	numBatches := int64(len(notes)) / batchSize
	if int64(len(notes))%batchSize != 0 {
		numBatches++
	}

	var end int64

	res := make([][]model.Note, 0, numBatches)

	for begin := int64(0); begin < int64(len(notes)); {
		end += batchSize
		if end > int64(len(notes)) {
			end = int64(len(notes))
		}

		res = append(res, notes[begin:end])
		begin += batchSize
	}

	return res, nil
}
