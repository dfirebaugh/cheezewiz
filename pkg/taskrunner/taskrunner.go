package taskrunner

import "time"

func Add(interval time.Duration, tickFunc func()) {
	// quietly check the resource status on an interval
	ticker := time.NewTicker(interval)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				go tickFunc()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
