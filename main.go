package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
)

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
	err := clone(repository)

	if err != nil {
		fmt.Println(err)
	}
}

func clone(repository string) error {
	destination := path.Join(os.Getenv("GOPATH"), "src", repository)
	fmt.Println(destination)
	err := os.MkdirAll(destination, 0777)

	if err != nil {
		return err
	}

	url := "git@" + strings.Replace(repository, "github.com/", "github.com:", 1) + ".git"
	cmd := exec.Command("git", "clone", "--depth", "1", url, destination)
	fmt.Println(cmd.Args)
	return cmd.Run()
}
