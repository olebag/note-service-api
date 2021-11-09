package main

import (
	"fmt"

	"github.com/scipie28/note-service-api/internal/app/api"
	"github.com/scipie28/note-service-api/internal/utills"
)

//123
func main() {

	fmt.Println(utills.SplitSlice([]int32{
		1, 4, 3},
		2))

	fmt.Println(utills.SwapKeyAndValue(map[int32]string{
		1: "one",
	}))

	fmt.Println(utills.FilterSlice([]string{
		"d", "r", "t", "b", "a"}))

	api.OpenCloseFile("text.txt")

	data := []api.User{
		{UserId: 1, ClassroomId: 23, DocumentId: 6},
		{UserId: 2, ClassroomId: 24, DocumentId: 7},
		{UserId: 3, ClassroomId: 23, DocumentId: 6},
		{UserId: 4, ClassroomId: 24, DocumentId: 7},
		{UserId: 5, ClassroomId: 23, DocumentId: 6},
		{UserId: 6, ClassroomId: 24, DocumentId: 7},
	}
	fmt.Println(utills.ConvertStructToMap(data))
	fmt.Println(utills.SplitSlizeUsers(data, 3))
}
