package main

import (
	"flag"
	"github.com/golang/glog"
	"monbus/pkg/busserver"
	"monbus/pkg/buswatcher"
	"sync"
	"time"
)

const (
	version = "v0.1.2"
)

func init() {
	flag.Parse()
}

func busWatcher(wg *sync.WaitGroup) {
	defer wg.Done()

	w, err := buswatcher.New()
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

func server(wg *sync.WaitGroup) {
	defer wg.Done()
	busserver.Test()
}

func main() {
	glog.Infof("Start monbus with version %s, time %v", version, time.Now())

	var wg sync.WaitGroup
	wg.Add(1)
	go busWatcher(&wg)
	go server(&wg)
	wg.Wait()
}
