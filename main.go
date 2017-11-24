package main

import "flag"
import "sync"
import "fmt"

func main() {
	flag.Parse()

	wg := sync.WaitGroup{}

	for _, arg := range flag.Args() {
		wg.Add(1)

		go func(repository string) {
			get(repository)
			wg.Done()
		}(arg)
	}

	wg.Wait()
}

func get(repository string) {
	fmt.Println(repository)
}
