package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var errCount int32
	var wg sync.WaitGroup
	ch := make(chan Task)

	if m <= 0 {
		m = len(tasks)
	}

	getTaskFromSlice := func() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, task := range tasks {
				ch <- task
				if atomic.LoadInt32(&errCount) >= int32(m) {
					break
				}
			}
			close(ch)
		}()
	}

	doingTask := func() {
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for task := range ch {
					if task() != nil {
						atomic.AddInt32(&errCount, 1)
					}
				}
			}()
		}
	}

	getTaskFromSlice()
	doingTask()
	wg.Wait()

	if errCount >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
