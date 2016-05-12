package retry

import (
	"sync"
	"time"
)

type CountRetryer struct {
	maxCount int
	interval time.Duration
	count    int
	started  bool
	m        sync.Mutex
}

func NewCountRetryer(maxCount int, interval time.Duration) *CountRetryer {
	return &CountRetryer{
		maxCount: maxCount,
		interval: interval,
	}
}

func (r *CountRetryer) Loop() bool {
	r.m.Lock()
	defer r.m.Unlock()
	if r.count == r.maxCount {
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

type DurationRetryer struct {
	limit    time.Time
	interval time.Duration
	finished bool
	started  bool
	m        sync.Mutex
}

func NewDurationRetryer(maxDuration, interval time.Duration) *DurationRetryer {
	return &DurationRetryer{
		limit:    time.Now().Add(maxDuration),
		interval: interval,
	}
}

func (r *DurationRetryer) Loop() bool {
	r.m.Lock()
	defer r.m.Unlock()
	if r.finished {
		return false
	}
	if r.started {
		time.Sleep(r.interval)
	} else {
		r.started = true
	}
	r.finished = time.Now().Add(r.interval).After(r.limit)
	return true
}
