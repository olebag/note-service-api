package api

import (
	"fmt"
	"os"
)

//Для своей предметной области описать структуры и
//реализовать методы работы с ними к примеру String()
//type User struct {
//	UserId uint64
//	// ...
//}

func OpenCloseFile(file string) {
	for i := 0; i < 5; i++ {
		data, err := os.Open(file)
		if err != nil {
			fmt.Println("Error", err)
			return
		}

		fmt.Println(data.Name())
		defer func(data *os.File) {
			err := data.Close()
			if err != nil {
				fmt.Println("Error", err)
			}
		}(data)
	}

}

type User struct {
	id          uint64
	UserId      uint32
	ClassroomId uint32
	DocumentId  uint32
}

func (a *User) String() {
	fmt.Println(a.id, a.UserId, a.ClassroomId, a.DocumentId)
}
