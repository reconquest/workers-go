# workers

That's really dead simple library that just reduces complexity of paralleling
processes in for-loops.

If you don't need to specify capacity of your thread pool than you just run all
routines like `go fn()`, but if you want to have some kind of capacity and do
not run more than N threads at the same time then this library is for you.


Example of usage:
```go
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
```

Output:
```
1s
1s
1s
2s
2s
2s
3s
3s
3s
4s
```

As you can see, we specified capacity as 3, it means that at the same time
could be running only 3 routines, while 3 routines are working, call `Run()` is
blocked, if it is not out of capacity then Run will start go-routine of your
function. This is why you have to run `Wait()` after for-loop.
