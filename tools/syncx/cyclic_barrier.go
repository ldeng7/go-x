package syncx

import (
	"context"
	"errors"
	"sync"
)

var ErrCyclicBarrierBroken = errors.New("broken barrier")

type cyclicBarrierRound struct {
	n      int
	w      chan struct{}
	b      chan struct{}
	broken bool
}

func newCyclicBarrierRound() *cyclicBarrierRound {
	return &cyclicBarrierRound{
		w: make(chan struct{}),
		b: make(chan struct{}),
	}
}

type CyclicBarrier struct {
	n  int
	cb func() error
	l  sync.RWMutex
	r  *cyclicBarrierRound
}

func NewCyclicBarrier(n int, onAllArrived func() error) *CyclicBarrier {
	if n <= 0 {
		panic("n must be positive")
	}
	return &CyclicBarrier{
		n:  n,
		cb: onAllArrived,
		l:  sync.RWMutex{},
		r:  newCyclicBarrierRound(),
	}
}

func (b *CyclicBarrier) Await(ctx context.Context) error {
	var ctxDoneCh <-chan struct{}
	if nil != ctx {
		ctxDoneCh = ctx.Done()
	}

	select {
	case <-ctxDoneCh:
		return ctx.Err()
	default:
	}

	b.l.Lock()
	if b.r.broken {
		b.l.Unlock()
		return ErrCyclicBarrierBroken
	}
	b.r.n++
	wCh, bCh, n := b.r.w, b.r.b, b.r.n
	b.l.Unlock()

	if n < b.n {
		select {
		case <-wCh:
			return nil
		case <-bCh:
			return ErrCyclicBarrierBroken
		case <-ctxDoneCh:
			b.doBreak(true)
			return ctx.Err()
		}
	} else if n == b.n {
		if nil != b.cb {
			if err := b.cb(); nil != err {
				b.doBreak(true)
				return err
			}
		}
		b.reset(true)
		return nil
	} else {
		panic("await called excessively")
	}
}

func (b *CyclicBarrier) doBreak(needLock bool) {
	if needLock {
		b.l.Lock()
		defer b.l.Unlock()
	}
	if !b.r.broken {
		b.r.broken = true
		close(b.r.b)
	}
}

func (b *CyclicBarrier) reset(safe bool) {
	b.l.Lock()
	defer b.l.Unlock()
	if safe {
		close(b.r.w)
	} else if b.r.n > 0 {
		b.doBreak(false)
	}
	b.r = newCyclicBarrierRound()
}

func (b *CyclicBarrier) Reset() {
	b.reset(false)
}

func (b *CyclicBarrier) NumWaiting() int {
	b.l.RLock()
	defer b.l.RUnlock()
	return b.r.n
}

func (b *CyclicBarrier) IsBroken() bool {
	b.l.RLock()
	defer b.l.RUnlock()
	return b.r.broken
}
