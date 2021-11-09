package utills

import (
	"fmt"
)

import "github.com/scipie28/note-service-api/internal/app/api"

var filter = []string{"a", "b", "c", "d"}

func SplitSlice(data []int32, batchSize int32) [][]int32 {
	if int32(len(data)) <= batchSize {
		return [][]int32{data}
	}

	numBatches := int32(len(data)) / batchSize
	if int32(len(data))%batchSize != 0 {
		numBatches++
	}

	var end int32

	res := make([][]int32, 0, numBatches)

	for begin := 0; begin < len(data); {
		end += batchSize
		if end > int32(len(data)) {
			end = int32(len(data))
		}

		res = append(res, data[begin:end])
		begin += int(batchSize)
	}

	return res
}

func SwapKeyAndValue(data map[int32]string) map[string]int32 {
	res := make(map[string]int32)

	for key, value := range data {
		if _, found := res[value]; found {
			fmt.Println("The key exists", value)
			continue
		}

		res[value] = key
	}

	return res
}

func FilterSlice(data []string) []string {
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

	return res
}

func ConvertStructToMap(user []api.User) (map[uint64]api.User, error) {
	res := make(map[uint64]api.User)
	for _, v := range user {
		res[uint64(v.UserId)] = v
	}
	return res, nil
}

func SplitSlizeUsers(user []api.User, butchSize uint32) [][]api.User {
	if uint32(len(user)) <= butchSize {
		return [][]api.User{}
	}

	numBatches := uint32(len(user)) / butchSize
	if uint32(len(user))%butchSize != 0 {
		numBatches++
	}

	var end uint32

	res := make([][]api.User, 0, numBatches)

	for begin := 0; begin < len(user); {
		end += butchSize
		if end > uint32(len(user)) {
			end = uint32(len(user))
		}

		res = append(res, user[begin:end])
		begin += int(butchSize)
	}
	return res
}
