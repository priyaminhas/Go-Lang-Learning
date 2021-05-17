package main

import (
	"fmt"
	"sync"
)

func main()  {
	counter := 0
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0 ; i< 100 ; i++{
		go func() {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			counter += 1
			fmt.Println(counter)
		}()
	}
	wg.Wait()
}
