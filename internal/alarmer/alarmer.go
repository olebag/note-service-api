package alarmer

//go:generate mockgen -destination=mocks/mock_alarmer.go -package=mocks . Alarmer

import (
	"time"
)

// Alarmer ...
type Alarmer interface {
	Alarm() <-chan struct{}
	Init() error
	Close()
}

type alarmer struct {
	duration time.Duration
	alarm    chan struct{}
	end      chan struct{}
}

func NewAlarmer(duration time.Duration) Alarmer {
	return &alarmer{
		duration: duration,
		alarm:    make(chan struct{}),
		end:      make(chan struct{}),
	}
}

func (a *alarmer) Init() error {
	go func() {
		ticker := time.NewTicker(a.duration)
		defer ticker.Stop()
		defer close(a.alarm)
		defer close(a.end)

		for {
			select {
			case <-ticker.C:
				a.alarm <- struct{}{}
			case <-a.end:
				return
			}
		}
	}()

	return nil
}

func (a *alarmer) Alarm() <-chan struct{} {
	return a.alarm
}

func (a *alarmer) Close() {
	a.end <- struct{}{}
}
