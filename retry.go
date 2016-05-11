package retry

import (
	"sync"
	"time"
)

type Retryer struct {
	maxTry   int
	interval time.Duration
	count    int
	started  bool
	m        sync.Mutex
}

func NewRetryer(maxTry int, interval time.Duration) *Retryer {
	return &Retryer{
		maxTry:   maxTry,
		interval: interval,
	}
}

func (r *Retryer) Loop() bool {
	r.m.Lock()
	defer r.m.Unlock()
	if r.count == r.maxTry {
		return false
	}
	r.count++
	if r.started {
		time.Sleep(r.interval)
	} else {
		r.started = true
	}
	return true
}
