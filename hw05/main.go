package main

import (
	"fmt"
)

func Run(tasks []func() error, N int, M int) error {
	if len(tasks) < N {
		N = len(tasks)
	}

	tasksCh := make(chan func() error, N)
	resCh := make(chan error, N)
	doneWorkGoroutineCh := make(chan struct{}, N)
	shutdown := make(chan string)

	//run task in N separate Goroutines
	for i := 1; i <= N; i++ {
		go func() {
			for task := range tasksCh {
				resCh <- task()
			}
			doneWorkGoroutineCh <- struct{}{}
		}()
	}

	//Send tasks to tasksCh and check results
	go func() {
		var suc int
		var err int
		var done bool
		var res string

		//Send first N tasks in to tasksCh
		for i := 0; i < N; i++ {
			tasksCh <- tasks[i]
		}

		addTaskIndex := N
		for resTask := range resCh {
			switch resTask.(type) {
			case error:
				err++
			default:
				suc++
			}
			if (err > 0 && err == M || suc+err == len(tasks)) && !done {
				done = true
				close(tasksCh)
				if err > 0 && err == M {
					res = fmt.Sprintln("Exit by error", "err", err, "suc", suc)
				} else {
					res = fmt.Sprintln("Exit by all done", "err", err, "suc", suc)
				}
			}
			if !done && addTaskIndex < len(tasks) {
				tasksCh <- tasks[addTaskIndex]
				addTaskIndex++
			}
		}
		shutdown <- res
	}()

	//Check all Goroutines done work
	go func() {
		var countDoneWorkGoroutines int
		for {
			<-doneWorkGoroutineCh
			countDoneWorkGoroutines++
			if countDoneWorkGoroutines == N {
				close(resCh)
				return
			}
		}
	}()

	fmt.Println(<-shutdown)
	return nil
}
