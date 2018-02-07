package workers

import (
	"sync"
)

// Workers is a group of processes that can be run in parallel mode in the same
// time but thread pool of Workers has maximum capacity.
type Workers struct {
	pool  chan struct{}
	group *sync.WaitGroup
}

// New creates instance of Workers with maximum capacity specified.
func New(capacity int) *Workers {
	return &Workers{
		pool:  make(chan struct{}, capacity),
		group: &sync.WaitGroup{},
	}
}

// Run given function in existing thread pool, if thread pool is out of
// capacity now, then function will be blocked until task from thread pool is
// done.
func (workers *Workers) Run(fn func()) {
	// mark one thread as running: write empty struct to buffered channel
	// the operation will be blocked if max count of workers already
	// running
	workers.pool <- struct{}{}

	workers.group.Add(1)

	go func() {
		defer workers.group.Done()
		defer func() {
			// free one item from buffered channel
			// defer is important here because fn() can panic.
			<-workers.pool
		}()

		fn()
	}()
}

// Wait for all processes to be done.
func (workers *Workers) Wait() {
	workers.group.Wait()
}
