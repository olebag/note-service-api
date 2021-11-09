package main

import (
	"fmt"
	"os"

	"github.com/scipie28/note-service-api/internal/app/api"
	"github.com/scipie28/note-service-api/internal/utills"
)

func main() {

	fmt.Println(utills.SplitSlice([]int32{
		1, 4, 3},
		2))

	fmt.Println(utills.SwapKeyAndValue(map[int32]string{
		1: "one",
	}))

	fmt.Println(utills.FilterSlice([]string{
		"d", "r", "t", "b", "a"}))

	fmt.Println(OpenCloseFile("text.txt"))

	data := []api.User{
		{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
		{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
		{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
		{Id: 5, UserId: 5, ClassroomId: 23, DocumentId: 6},
		{Id: 555, UserId: 6, ClassroomId: 24, DocumentId: 7},
	}

	fmt.Println(utills.ConvertSliceToMap(data))

	fmt.Println(utills.SplitSliceUsers(data, 3))
	data[0].String()
}

func OpenCloseFile(file string) error {
	for i := 0; i < 5; i++ {
		data, err := os.Open(file)

		defer func(data *os.File) {
			err := data.Close()
			if err != nil {
				fmt.Printf("Error closing file %s", err)
			}
		}(data)

		if err != nil {
			return err
		}

		fmt.Println(data.Name())
	}

	return nil
}
