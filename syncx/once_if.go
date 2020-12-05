package syncx

import (
	"sync"
	"sync/atomic"
)

type OnceIf struct {
	m    sync.Mutex
	done uint32
}

func (o *OnceIf) Done() bool {
	return atomic.LoadUint32(&o.done) == 1
}

func (o *OnceIf) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	return o.doSlow(f)
}

func (o *OnceIf) doSlow(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 {
		if err = f(); nil == err {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}
