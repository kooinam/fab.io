package models

import (
	"time"
)

type Runner struct {
	handler   func()
	interval  time.Duration
	elapsed   time.Duration
	shouldRun bool
	Ch        chan int
}

// MakeRunner used to instantiate runner instance
func MakeRunner(handler func(), interval time.Duration) *Runner {
	runner := &Runner{
		handler:   handler,
		interval:  interval,
		Ch:        make(chan int),
		shouldRun: false,
	}

	go func() {
		status := <-runner.Ch

		if status == 0 {
			if runner.shouldRun {
				runner.shouldRun = false
			}
		} else if status == 1 {
			if !runner.shouldRun {
				runner.shouldRun = true

				runner.start()
			}
		}
	}()

	return runner
}

func (runner *Runner) start() {
	go func() {
		for {
			if runner.shouldRun {
				runner.run()
			} else {
				break
			}
		}
	}()
}

func (runner *Runner) run() {
	runner.handler()

	time.Sleep(runner.interval * time.Second)

	runner.elapsed += runner.interval * time.Second
}
