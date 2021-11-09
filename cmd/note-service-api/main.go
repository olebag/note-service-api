package main

import (
	"fmt"
	"os"

	"github.com/scipie28/note-service-api/internal/app/api"
	"github.com/scipie28/note-service-api/internal/utills"
)

func main() {

	fmt.Println(utills.SwapKeyAndValue(map[int32]string{1: "one"}))

	fmt.Println(utills.FilterSlice([]string{"d", "r", "t", "b", "a"}))

	res := OpenCloseFile("cmd/note-service-api/text.txt")
	if res != nil {
		fmt.Printf("OpenCloseFile() function execution error %s\n\n", res)
	} else {
		fmt.Printf("OpenCloseFile() function comleted\n\n")
	}

	data := []api.User{
		{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
		{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
		{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
		{Id: 5, UserId: 5, ClassroomId: 23, DocumentId: 6},
		{Id: 555, UserId: 6, ClassroomId: 24, DocumentId: 7},
	}

	fmt.Println(utills.ConvertSliceToMap(data))

	fmt.Println(utills.SplitSlice(data, 3))

	data[0].String()
}

func OpenCloseFile(file string) error {
	for i := 1; i < 5; i++ {
		data, err := os.Open(file)

		f := func() {
			defer func(data *os.File) {
				err := data.Close()
				if err != nil {
					fmt.Printf("Failed to closing file: %s", err)
				}
				fmt.Printf("Закрыл файл в %v раз.\n\n", i)
			}(data)
		}

		if err != nil {
			return err
		}

		fmt.Printf("Открыл файл в %v раз. \n", i)
		f()
	}

	return nil
}
