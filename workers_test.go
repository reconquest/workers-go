package workers_test

import (
	"fmt"
	"time"

	"github.com/reconquest/workers-go"
)

func ExampleWorkers() {
	start := time.Now().Truncate(time.Second)

	pool := workers.New(3)
	for i := 0; i < 10; i++ {
		pool.Run(
			func() {
				time.Sleep(time.Second)

				fmt.Printf(
					"%v\n",
					time.Since(start).Truncate(time.Second),
				)
			},
		)
	}

	pool.Wait()

	// Output:
	// 1s
	// 1s
	// 1s
	// 2s
	// 2s
	// 2s
	// 3s
	// 3s
	// 3s
	// 4s
}
