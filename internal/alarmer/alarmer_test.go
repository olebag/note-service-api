package alarmer

import (
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

func TestAlarmer(t *testing.T) {
	maxRes := 6
	minRes := 4
	num := 0

	alarmerTest, errNewAlarm := NewAlarmer(1 * time.Second)
	if errNewAlarm != nil {
		log.Printf("failed to crating new alarmer")
	}

	errInitAlarm := alarmerTest.Init()
	if errInitAlarm != nil {
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

	require.Nil(t, errInitAlarm)
	require.Nil(t, errNewAlarm)
	require.True(t, success)
}
