package utills

import (
	"fmt"

	"github.com/scipie28/note-service-api/internal/app/api"
)

var filter = []string{"a", "b", "c", "d"}

func SwapKeyAndValue(data map[int32]string) (map[string]int32, error) {
	res := make(map[string]int32)

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

func SplitSlice(users []api.Note, butchSize uint32) ([][]api.Note, error) {
	if uint32(len(users)) <= butchSize {
		return [][]api.Note{}, nil
	}

	numBatches := uint32(len(users)) / butchSize
	if uint32(len(users))%butchSize != 0 {
		numBatches++
	}

	var end uint32

	res := make([][]api.Note, 0, numBatches)

	for begin := uint32(0); begin < uint32(len(users)); {
		end += butchSize
		if end > uint32(len(users)) {
			end = uint32(len(users))
		}

		res = append(res, users[begin:end])
		begin += butchSize
	}

	return res, nil
}
