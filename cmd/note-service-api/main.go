package main

import (
	"fmt"
	"os"

	"github.com/scipie28/note-service-api/internal/app/api"
	"github.com/scipie28/note-service-api/internal/utills"
)

func main() {
	val1, err1 := utills.FilterSlice([]string{"d", "r", "t", "b", "a"})
	if err1 != nil {
		fmt.Printf("error to start function FilterSlice %s", err1)
	} else {
		fmt.Println(val1)
	}

	val2, err2 := utills.SwapKeyAndValue(map[int32]string{1: "one"})
	if err2 != nil {
		fmt.Printf("error to start function SwapKeyAndValue %s", err2)
	} else {
		fmt.Println(val2)
	}

	err3 := OpenCloseFile("cmd/note-service-api/text.txt")
	if err3 != nil {
		fmt.Printf("OpenCloseFile() function execution error %s\n\n", err3)
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

	val4, err4 := utills.ConvertSliceToMap(data)
	if err4 != nil {
		fmt.Printf("error to start function ConvertSliceToMap %s", err4)
	} else {
		fmt.Println(val4)
	}

	val5, err5 := utills.SplitSlice(data, 3)
	if err5 != nil {
		fmt.Printf("error to start function SplitSlice %s", err5)
	} else {
		fmt.Println(val5)
	}

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
