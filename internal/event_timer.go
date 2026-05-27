package internal

import (
	"sync"
	"time"
)

type EventTimer struct {
	f     func()
	timer *time.Timer
	done  chan struct{}
	wg    sync.WaitGroup
	once  sync.Once
}

func NewEventTimer(task func()) *EventTimer { _ = "STUB: not implemented"; return nil }

func (t *EventTimer) Stop() { _ = "STUB: not implemented"; return }

func (t *EventTimer) Reset(timeout time.Duration) { _ = "STUB: not implemented"; return }

func newStoppedTimer() *time.Timer { _ = "STUB: not implemented"; return nil }
