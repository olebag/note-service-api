package main

import (
	"fmt"
	"os"

	"github.com/scipie28/note-service-api/internal/app/api"
	"github.com/scipie28/note-service-api/internal/utills"
)

func main() {
	var err error

	resFilter, err := utills.FilterSlice([]string{"d", "r", "t", "b", "a"})
	if err != nil {
		fmt.Printf("error to start function FilterSlice %s", err)
		return
	}

	fmt.Println(resFilter)

	resSwap, err := utills.SwapKeyAndValue(map[int32]string{1: "one"})
	if err != nil {
		fmt.Printf("error to start function SwapKeyAndValue %s", err)
		return
	}

	fmt.Println(resSwap)

	err = OpenCloseFile("cmd/note-service-api/text.txt")
	if err != nil {
		fmt.Printf("OpenCloseFile() function execution error %s\n\n", err)
	}

	fmt.Printf("OpenCloseFile() function comleted\n\n")

	data := []api.User{
		{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
		{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
		{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
		{Id: 5, UserId: 5, ClassroomId: 23, DocumentId: 6},
		{Id: 555, UserId: 6, ClassroomId: 24, DocumentId: 7},
	}

	resConvert, err := utills.ConvertSliceToMap(data)
	if err != nil {
		fmt.Printf("error to start function ConvertSliceToMap %s", err)
	}

	fmt.Println(resConvert)

	resSplit, err := utills.SplitSlice(data, 3)
	if err != nil {
		fmt.Printf("error to start function SplitSlice %s", err)
	}

	fmt.Println(resSplit)

	for _, v := range data {
		v.String()
	}
}

func OpenCloseFile(file string) error {
	for i := 1; i < 5; i++ {
		err := func() error {
			data, err := os.Open(file)

			defer func(data *os.File) {
				err = data.Close()
				if err != nil {
					fmt.Printf("failed to closing file: %s", err)
					return
				}

				fmt.Printf("Closed the file %v times.\n\n", i)
			}(data)

			if err != nil {
				return err
			}

			fmt.Printf("Opened the file %v times. \n", i)

			return nil
		}()

		if err != nil {
			return err
		}
	}

	return nil
}
