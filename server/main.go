package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/facebookgo/inject"
)

//TODO: Logging
//TODO: Config

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	run(func() { wg.Wait() }, func() { wg.Done() })
}

func run(start, stop func()) {
	term := make(chan os.Signal)
	signal.Notify(term, syscall.SIGINT)
	signal.Notify(term, syscall.SIGKILL)
	unhandled := make(chan struct{})

	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer func() {
			wg.Done()
			if r := recover(); r != nil {
				//logger.HadPanic("Server Shutting down on panic", r)
				close(unhandled)
			}
		}()
		//logger.Inform("Running!")
		start()
	}()

	select {
	case <-unhandled:
	//	logger.Inform("Got unhandled exception signal")
	case <-term:
		//	logger.Inform("Got shutdown signal")
	}
	//logger.Inform("Stopping... ")
	stop()
	//logger.Inform("Waiting on server to stop...")
	wg.Wait()
	//logger.Inform("Done!")
}

// TODO: pass a list of Key Values
func dependencies() {
	var graph inject.Graph
	err := graph.Provide(
		&inject.Object{Value: "", Name: "myname"},
	)
	if err != nil {
		panic(err)
	}

	if err := graph.Populate(); err != nil {
		panic(err)
	}
}
