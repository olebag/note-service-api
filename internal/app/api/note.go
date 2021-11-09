package api

import (
	"fmt"
)

// User...
type User struct {
	Id          uint64
	UserId      uint32
	ClassroomId uint32
	DocumentId  uint32
}

func (u *User) String() {
	fmt.Printf(
		"Id = %v; UserId = %v; ClasssroomId = %v; DocumentId = %v\n",
		u.Id, u.UserId, u.ClassroomId, u.DocumentId)
}
