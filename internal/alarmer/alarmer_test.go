package alarmer

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAlarmer(t *testing.T) {
	t.Run("duration value more zero", func(t *testing.T) {
		maxRes := 6
		minRes := 4
		num := 0

		alarmerTest, err := NewAlarmer(1 * time.Second)
		require.NoError(t, err)

		err = alarmerTest.Init()
		require.NoError(t, err)

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

	t.Run("duration value equal zero", func(t *testing.T) {
		expectedErr := "invalid duration"

		_, err := NewAlarmer(0 * time.Second)
		require.NotNil(t, err)
		require.Equal(t, expectedErr, err.Error())
	})

	t.Run("duration value less zero", func(t *testing.T) {
		expectedErr := "invalid duration"

		_, err := NewAlarmer(-5 * time.Second)
		require.NotNil(t, err)
		require.Equal(t, expectedErr, err.Error())
	})
}
