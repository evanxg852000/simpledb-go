package utils

import (
	"sync"
	"time"
)

func WaitCondWithTimeout(cond *sync.Cond, timeout time.Duration) {
	go func(cond *sync.Cond, timeout time.Duration) {
		time.Sleep(timeout)
		cond.Broadcast()
	}(cond, timeout)
	cond.Wait()
}
