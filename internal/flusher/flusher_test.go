package flusher

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/scipie28/note-service-api/internal/app/api"
	moksRepo "github.com/scipie28/note-service-api/internal/repo/mocks"
	"github.com/stretchr/testify/require"
)

func TestFlusher_Flush(t *testing.T) {
	var (
		mockCtrl = gomock.NewController(t)
	)

	mokNoteRepo := moksRepo.NewMockRepo(mockCtrl)
	noteFlusher := New(mokNoteRepo)

	t.Run("success case", func(t *testing.T) {
		mokNoteRepo.EXPECT().MultiAdd([]api.Note{
			{Id: 1,
				UserId:      2,
				ClassroomId: 3,
				DocumentId:  4},
		}).Return(int32(0), nil).Times(1)
		mokNoteRepo.EXPECT().MultiAdd([]api.Note{
			{Id: 5,
				UserId:      6,
				ClassroomId: 7,
				DocumentId:  8},
		}).Return(int32(0), nil).Times(1)

		req := []api.Note{
			{Id: 1,
				UserId:      2,
				ClassroomId: 3,
				DocumentId:  4},
			{Id: 5,
				UserId:      6,
				ClassroomId: 7,
				DocumentId:  8},
		}

		res, err := noteFlusher.Flush(req, 1)

		var expectedRes []api.Note

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})
}
