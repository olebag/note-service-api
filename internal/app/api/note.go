package api

import (
	"fmt"
)

// Note...
type Note struct {
	Id          uint64
	UserId      uint32
	ClassroomId uint32
	DocumentId  uint32
}

func (n *Note) String() {
	fmt.Printf(
		"Id = %v; UserId = %v; ClasssroomId = %v; DocumentId = %v\n",
		n.Id, n.UserId, n.ClassroomId, n.DocumentId)
}
