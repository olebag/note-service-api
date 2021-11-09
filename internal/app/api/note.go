package api

import (
	"fmt"
)

// User
type User struct {
	Id          uint64
	UserId      uint32
	ClassroomId uint32
	DocumentId  uint32
}

func (u *User) String() {
	fmt.Printf(
		" Id = %v\n UserId = %v\n ClasssroomId = %v\n DocumentId = %v",
		u.Id, u.UserId, u.ClassroomId, u.DocumentId)
}
