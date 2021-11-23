package saver

import (
	"log"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/scipie28/note-service-api/internal/alarmer"
	"github.com/scipie28/note-service-api/internal/app/api"
	"github.com/scipie28/note-service-api/internal/flusher"
	mocksRepo "github.com/scipie28/note-service-api/internal/repo/mocks"
	"github.com/stretchr/testify/require"
)

func TestSaver(t *testing.T) {
	var (
		mockCtrl        = gomock.NewController(t)
		mockNoteRepo    = mocksRepo.NewMockRepo(mockCtrl)
		lossAllDataMode = true
	)

	noteFlusher := flusher.NewFlusher(mockNoteRepo)

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

	t.Run("input capacity equal zero", func(t *testing.T) {
		expectedError := "error input value: capacity"

		alarmerTest, _ := alarmer.NewAlarmer(20 * time.Millisecond)

		_, err := NewSaver(0, 3, noteFlusher, alarmerTest, lossAllDataMode)
		require.NotNil(t, err)
		require.Equal(t, expectedError, err.Error())
	})

	t.Run("input batch size equal zero", func(t *testing.T) {
		expectedError := "error input value: batch size"

		alarmerTest, _ := alarmer.NewAlarmer(20 * time.Millisecond)

		_, err := NewSaver(3, 0, noteFlusher, alarmerTest, lossAllDataMode)
		require.NotNil(t, err)
		require.Equal(t, expectedError, err.Error())
	})

	t.Run("time alarmer less time write in buffer", func(t *testing.T) {

		t.Run("input capacity equal one", func(t *testing.T) {
			mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(8)
			wg := sync.WaitGroup{}
			wg.Add(8)
			alarmerTest, _ := alarmer.NewAlarmer(5 * time.Millisecond)

			saverTest, err := NewSaver(1, 3, noteFlusher, alarmerTest, lossAllDataMode)
			if err != nil {
				log.Printf("fail to crate new saver: %s", err.Error())
			}

			err = saverTest.Init()
			if err != nil {
				log.Printf("failed to initialized saver: %s", err.Error())
				return
			}
			defer saverTest.Close()

			for _, val := range req {
				errReq := saverTest.Save(val)
				if errReq != nil {
					log.Printf("failed to save %s", errReq.Error())
				}
				time.Sleep(20 * time.Millisecond)

			}

			require.Nil(t, err)
		})

		t.Run("input capacity less slice notes", func(t *testing.T) {
			mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(8)

			alarmerTest, _ := alarmer.NewAlarmer(5 * time.Millisecond)

			saverTest, err := NewSaver(2, 3, noteFlusher, alarmerTest, lossAllDataMode)
			if err != nil {
				log.Printf("fail to crate new saver: %s", err.Error())
			}

			err = saverTest.Init()
			if err != nil {
				log.Printf("failed to initialized saver: %s", err.Error())
				return
			}
			defer saverTest.Close()

			for _, val := range req {
				errReq := saverTest.Save(val)
				if errReq != nil {
					log.Printf("failed to save: %s", errReq.Error())
				}
				time.Sleep(20 * time.Millisecond)
			}

			require.Nil(t, err)
		})

		t.Run("input capacity more slice notes", func(t *testing.T) {
			mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(8)

			alarmerTest, _ := alarmer.NewAlarmer(5 * time.Millisecond)

			saverTest, err := NewSaver(9, 3, noteFlusher, alarmerTest, lossAllDataMode)
			if err != nil {
				log.Printf("fail to crate new saver: %s", err.Error())
			}

			err = saverTest.Init()
			if err != nil {
				log.Printf("failed to initialized saver: %s", err.Error())
				return
			}
			defer saverTest.Close()

			for _, val := range req {
				errReq := saverTest.Save(val)
				if errReq != nil {
					log.Printf("failed to save: %s", errReq.Error())
				}
				time.Sleep(20 * time.Millisecond)
			}

			require.Nil(t, err)
		})
	})

	t.Run("time write in buffer less time alarmer", func(t *testing.T) {
		t.Run("input capacity equal one", func(t *testing.T) {
			mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(2)

			alarmerTest, _ := alarmer.NewAlarmer(20 * time.Millisecond)

			saverTest, err := NewSaver(1, 3, noteFlusher, alarmerTest, lossAllDataMode)
			if err != nil {
				log.Printf("fail to crate new saver: %s", err.Error())
			}

			err = saverTest.Init()
			if err != nil {
				log.Printf("failed to initialized saver: %s", err.Error())
				return
			}
			defer saverTest.Close()

			for _, val := range req {
				errReq := saverTest.Save(val)
				if errReq != nil {
					log.Printf("failed to save: %s", errReq.Error())
				}
				time.Sleep(5 * time.Millisecond)
			}

			require.Nil(t, err)
		})

		t.Run("capacity less slice notes", func(t *testing.T) {
			t.Run("input batch size equal capacity", func(t *testing.T) {
				mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(2)

				alarmerTest, _ := alarmer.NewAlarmer(20 * time.Millisecond)

				saverTest, err := NewSaver(6, 6, noteFlusher, alarmerTest, lossAllDataMode)
				if err != nil {
					log.Printf("fail to crate new saver: %s", err.Error())
				}

				err = saverTest.Init()
				if err != nil {
					log.Printf("failed to initialized saver: %s", err.Error())
					return
				}
				defer saverTest.Close()

				for _, val := range req {
					err := saverTest.Save(val)
					if err != nil {
						log.Printf("failed to save: %s", err.Error())
					}
					time.Sleep(5 * time.Millisecond)
				}

				require.Nil(t, err)
			})

			t.Run("input batch size less capacity", func(t *testing.T) {
				mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(4)

				alarmerTest, _ := alarmer.NewAlarmer(20 * time.Millisecond)

				saverTest, err := NewSaver(6, 2, noteFlusher, alarmerTest, lossAllDataMode)
				if err != nil {
					log.Printf("fail to crate new saver: %s", err.Error())
				}

				err = saverTest.Init()
				if err != nil {
					log.Printf("failed to initialized saver: %s", err.Error())
					return
				}
				defer saverTest.Close()

				for _, val := range req {
					errReq := saverTest.Save(val)
					if errReq != nil {
						log.Printf("failed to save: %s", errReq.Error())
					}
					time.Sleep(5 * time.Millisecond)
				}

				require.Nil(t, err)
			})

			t.Run("input batch size more capacity", func(t *testing.T) {
				mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(2)

				alarmerTest, _ := alarmer.NewAlarmer(20 * time.Millisecond)

				saverTest, err := NewSaver(2, 6, noteFlusher, alarmerTest, lossAllDataMode)
				if err != nil {
					log.Printf("fail to crate new saver: %s", err.Error())
				}

				err = saverTest.Init()
				if err != nil {
					log.Printf("failed to initialized saver: %s", err.Error())
					return
				}
				defer saverTest.Close()

				for _, val := range req {
					err := saverTest.Save(val)
					if err != nil {
						log.Printf("failed to save: %s", err.Error())
					}
					time.Sleep(5 * time.Millisecond)
				}

				require.Nil(t, err)
			})
		})

		t.Run("capacity more slice notes", func(t *testing.T) {
			t.Run("input batch size equal capacity", func(t *testing.T) {
				mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(1)

				alarmerTest, _ := alarmer.NewAlarmer(40 * time.Millisecond)

				saverTest, err := NewSaver(9, 9, noteFlusher, alarmerTest, lossAllDataMode)
				if err != nil {
					log.Printf("fail to crate new saver: %s", err.Error())
				}

				err = saverTest.Init()
				if err != nil {
					log.Printf("failed to initialized saver: %s", err.Error())
					return
				}
				defer saverTest.Close()

				for _, val := range req {
					err := saverTest.Save(val)
					if err != nil {
						log.Printf("failed to save: %s", err.Error())
					}
					time.Sleep(5 * time.Millisecond)
				}

				require.Nil(t, err)
			})

			t.Run("input batch size less capacity", func(t *testing.T) {
				mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(2)

				alarmerTest, _ := alarmer.NewAlarmer(40 * time.Millisecond)

				saverTest, err := NewSaver(9, 4, noteFlusher, alarmerTest, lossAllDataMode)
				if err != nil {
					log.Printf("fail to crate new saver: %s", err.Error())
				}

				err = saverTest.Init()
				if err != nil {
					log.Printf("failed to initialized saver: %s", err.Error())
					return
				}
				defer saverTest.Close()

				for _, val := range req {
					err := saverTest.Save(val)
					if err != nil {
						log.Printf("failed to save: %s", err.Error())
					}
					time.Sleep(5 * time.Millisecond)
				}

				require.Nil(t, err)
			})

			t.Run("input batch size more capacity", func(t *testing.T) {
				mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(1)

				alarmerTest, _ := alarmer.NewAlarmer(40 * time.Millisecond)

				saverTest, err := NewSaver(9, 14, noteFlusher, alarmerTest, lossAllDataMode)
				if err != nil {
					log.Printf("fail to crate new saver: %s", err.Error())
				}

				err = saverTest.Init()
				if err != nil {
					log.Printf("failed to initialized saver: %s", err.Error())
					return
				}
				defer saverTest.Close()

				for _, val := range req {
					err := saverTest.Save(val)
					if err != nil {
						log.Printf("failed to save: %s", err.Error())
					}
					time.Sleep(5 * time.Millisecond)
				}

				require.Nil(t, err)
			})
		})
	})

	//t.Run("time alarmer equal time write in buffer", func(t *testing.T) {
	//	t.Run("input capacity equal one", func(t *testing.T) {
	//		mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(8)
	//
	//		alarmerTest, _ := alarmer.NewAlarmer(10 * time.Millisecond)
	//
	//		saverTest, err := NewSaver(1, 3, noteFlusher, alarmerTest, lossAllDataMode)
	//		if err != nil {
	//			log.Printf("failed to crate new saver: %s", err.Error())
	//		}
	//
	//		err = saverTest.Init()
	//		if err != nil {
	//			log.Printf("failed to initialized saver: %s", err.Error())
	//			return
	//		}
	//		defer saverTest.Close()
	//
	//		for _, val := range req {
	//			err := saverTest.Save(val)
	//			if err != nil {
	//				log.Printf("failed to save: %s", err.Error())
	//			}
	//			time.Sleep(10 * time.Millisecond)
	//		}
	//
	//		require.Nil(t, err)
	//	})
	//
	//	t.Run("input capacity more slice notes", func(t *testing.T) {
	//		mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(8)
	//
	//		alarmerTest, _ := alarmer.NewAlarmer(10 * time.Millisecond)
	//
	//		saverTest, err := NewSaver(9, 6, noteFlusher, alarmerTest, lossAllDataMode)
	//		if err != nil {
	//			log.Printf("fail to crate new saver: %s", err.Error())
	//		}
	//
	//		err = saverTest.Init()
	//		if err != nil {
	//			log.Printf("failed to initialized saver: %s", err.Error())
	//			return
	//		}
	//		defer saverTest.Close()
	//
	//		for _, val := range req {
	//			err := saverTest.Save(val)
	//			if err != nil {
	//				log.Printf("failed to save: %s", err.Error())
	//			}
	//			time.Sleep(10 * time.Millisecond)
	//		}
	//
	//		require.Nil(t, err)
	//	})
	//
	//	t.Run("input capacity less slice notes", func(t *testing.T) {
	//		mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(8)
	//
	//		alarmerTest, _ := alarmer.NewAlarmer(10 * time.Millisecond)
	//
	//		saverTest, err := NewSaver(3, 2, noteFlusher, alarmerTest, lossAllDataMode)
	//		if err != nil {
	//			log.Printf("failed to crate new saver: %s", err.Error())
	//		}
	//
	//		err = saverTest.Init()
	//		if err != nil {
	//			log.Printf("failed to initialized saver: %s", err.Error())
	//			return
	//		}
	//		defer saverTest.Close()
	//
	//		for _, val := range req {
	//			err := saverTest.Save(val)
	//			if err != nil {
	//				log.Printf("failed to save: %s", err.Error())
	//			}
	//			time.Sleep(10 * time.Millisecond)
	//		}
	//
	//		require.Nil(t, err)
	//	})
	//
	//})

}
