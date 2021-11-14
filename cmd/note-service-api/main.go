package main

import (
	"fmt"
	"os"

	"github.com/scipie28/note-service-api/internal/app/api"
	"github.com/scipie28/note-service-api/internal/flusher"
	"github.com/scipie28/note-service-api/internal/repo"
	"github.com/scipie28/note-service-api/internal/utills"
)

func main() {
	var err error

	dataMock := []api.Note{
		{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
		{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
		{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
		{Id: 5, UserId: 5, ClassroomId: 23, DocumentId: 6},
		{Id: 555, UserId: 6, ClassroomId: 24, DocumentId: 7},
	}

	noteRepo := repo.New(0)
	noteFlusher := flusher.New(noteRepo)

	flush, err := noteFlusher.Flush(dataMock, 2)
	if err != nil {
		fmt.Println("Программа внезапно прервалась", err)
		fmt.Println("Не успел передать:", flush)

		return
	}

	filteredSlice, err := utills.FilterSlice([]string{"d", "r", "t", "b", "a"})
	if err != nil {
		fmt.Printf("failed to filtering slice: %s", err)
		return
	}

	fmt.Println(filteredSlice)

	//---------------------------
	swappedMap, err := utills.SwapKeyAndValue(map[int32]string{1: "one"})
	if err != nil {
		fmt.Printf("error to start function SwapKeyAndValue %s", err)
		return
	}

	fmt.Println(swappedMap)

	err = OpenCloseFile("cmd/note-service-api/text.txt")
	if err != nil {
		fmt.Printf("OpenCloseFile() function execution error %s\n\n", err)
	}

	fmt.Printf("OpenCloseFile() function comleted\n\n")

	data := []api.Note{
		{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
		{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
		{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
		{Id: 5, UserId: 5, ClassroomId: 23, DocumentId: 6},
		{Id: 555, UserId: 6, ClassroomId: 24, DocumentId: 7},
	}

	dataMap, err := utills.ConvertSliceToMap(data)
	if err != nil {
		fmt.Printf("error to start function ConvertSliceToMap %s", err)
	}

	fmt.Println(dataMap)

	splitSlice, err := utills.SplitSlice(data, 5)
	if err != nil {
		fmt.Printf("error to start function SplitSlice %s", err)
	}

	fmt.Println(splitSlice)

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
