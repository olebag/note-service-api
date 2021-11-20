package saver

import (
	"log"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/scipie28/note-service-api/internal/alarmer"
	"github.com/scipie28/note-service-api/internal/app/api"
	"github.com/scipie28/note-service-api/internal/flusher"
	mocksRepo "github.com/scipie28/note-service-api/internal/repo/mocks"
	"github.com/stretchr/testify/require"
)

func TestSaver_Save(t *testing.T) {
	var (
		mockCtrl = gomock.NewController(t)
	)
	var mockNoteRepo = mocksRepo.NewMockRepo(mockCtrl)
	gomock.InOrder(
		mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(1),
		mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(1),
		mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(1),
		mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(1),
		mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(1),
		mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(1),
		mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(1),
		mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(1),
	)

	req := []api.Note{
		{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
		{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
		{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
		{Id: 11, UserId: 11, ClassroomId: 123, DocumentId: 16},
		{Id: 21, UserId: 21, ClassroomId: 124, DocumentId: 17},
		{Id: 31, UserId: 31, ClassroomId: 123, DocumentId: 16},
		{Id: 41, UserId: 41, ClassroomId: 124, DocumentId: 17},
	}

	noteFlusher := flusher.NewFlusher(mockNoteRepo)
	exmAlarmer := alarmer.NewAlarmer(1000 * time.Millisecond)
	exmSaver := NewSaver(1, 1, noteFlusher, exmAlarmer, true)

	errAlarm := exmAlarmer.Init()
	if errAlarm != nil {
		log.Printf("failed to initialized alarm: %s", errAlarm.Error())
		return
	}

	errSaver := exmSaver.Init()
	if errAlarm != nil {
		log.Printf("failed to initialized Saver: %s", errSaver.Error())
		return
	}

	for _, val := range req {
		defer exmSaver.Close()

		err := exmSaver.Save(val)
		if err != nil {
			log.Printf("failed to save %s", err.Error())
		}
		time.Sleep(1000 * time.Millisecond)
	}

	require.Nil(t, errSaver)
	require.Nil(t, errAlarm)
}
