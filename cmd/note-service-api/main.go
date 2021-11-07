package main

import (
	"fmt"

	"github.com/scipie28/note-service-api/internal/utills"
)

func main() {

	fmt.Println(utills.SplitSlice([]int32{
		1, 4, 3},
		2))

	fmt.Println(utills.SwapKeyAndValue(map[int32]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "five",
		8: "five",
		9: "five",
	}))

	fmt.Println(utills.FilterSlice([]string{
		"d", "r", "t", "b", "a"}))
}
