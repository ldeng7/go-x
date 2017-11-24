package baffle

import (
	"sync"
	"time"
)

type Baffle struct {
	intl  int64
	thres uint
	dur   int64

	lock     *sync.Mutex
	tick     int64
	count    uint
	recovery int64
}

func Init(interval int64, threshold uint, fuseDuration int64) *Baffle {
	b := &Baffle{
		intl:  interval,
		thres: threshold,
		dur:   fuseDuration,
		lock:  &sync.Mutex{},
	}
	return b
}

func (b *Baffle) CheckIn() bool {
	if b.recovery == 0 {
		return true
	}
	if time.Now().Unix() >= b.recovery {
		b.recovery = 0
		return true
	}
	return false
}

func (b *Baffle) Hit() {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.recovery != 0 {
		return
	}
	now := time.Now().Unix()
	if now >= b.tick {
		b.count = 0
		b.tick = now + b.intl
	}
	b.count += 1
	if b.count >= b.thres {
		b.recovery = now + b.dur
	}
}
