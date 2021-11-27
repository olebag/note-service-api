package note_v1

import (
	"github.com/scipie28/note-service-api/internal/app/service/note"
	desc "github.com/scipie28/note-service-api/pkg/note_v1"
)

type Note struct {
	desc.UnimplementedNoteV1Server

	noteService note.INote
}

