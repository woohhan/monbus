package main

import (
	"flag"
	"github.com/golang/glog"
	"monbus/pkg/busserver"
	"monbus/pkg/buswatcher"
	"monbus/pkg/storage"
	"sync"
	"time"
)

const (
	version = "v0.1.2"
)

func init() {
	flag.Parse()
}

func busWatcher(s *storage.Storage, wg *sync.WaitGroup) {
	defer wg.Done()

	w, err := buswatcher.New(s)
	if err != nil {
		panic(err)
	}
	defer func() {
		w.Close()
	}()

	w.AddStation(4, "미금역")
	w.AddStation(45, "주엽역(중)")
	w.Watch(10*time.Second, 30*time.Minute)
}

func server(s *storage.Storage, wg *sync.WaitGroup) {
	defer wg.Done()
	b := busserver.New(s)
	b.Run()
}

func main() {
	glog.Infof("Start monbus with version %s, time %v", version, time.Now())

	// init storage
	s, err := storage.New()
	if err != nil {
		panic(err)
	}

	// run watcher, server
	var wg sync.WaitGroup
	wg.Add(1)
	go busWatcher(s, &wg)
	go server(s, &wg)
	wg.Wait()
}
