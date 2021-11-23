package alarmer

import (
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

func TestAlarmer(t *testing.T) {
	t.Run("input duration more zero", func(t *testing.T) {
		maxRes := 6
		minRes := 4
		num := 0

		alarmerTest, err := NewAlarmer(1 * time.Second)
		if err != nil {
			log.Printf("failed to crating new alarmer")
		}

		err = alarmerTest.Init()
		if err != nil {
			log.Printf("failed to initialize alarmer")
		}

		timer := time.NewTimer(5 * time.Second)
		wg := sync.WaitGroup{}

		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-alarmerTest.Alarm():
					num++
				case <-timer.C:
					return
				}
			}
		}()
		wg.Wait()

		success := false
		if (num >= minRes) && (num <= maxRes) {
			success = true
		}

		require.Nil(t, err)
		require.True(t, success)
	})

	t.Run("input duration equal zero", func(t *testing.T) {
		expectedErr := "error input value: duration"

		_, err := NewAlarmer(0 * time.Second)

		require.NotNil(t, err)
		require.Equal(t, expectedErr, err.Error())
	})

	t.Run("input duration less zero", func(t *testing.T) {
		expectedErr := "error input value: duration"

		_, err := NewAlarmer(-5 * time.Second)

		require.NotNil(t, err)
		require.Equal(t, expectedErr, err.Error())
	})
}
