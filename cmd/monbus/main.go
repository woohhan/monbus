package main

import (
	"flag"
	"github.com/golang/glog"
	"monbus/pkg/watcher"
	"time"
)

const (
	version = "v0.1.2"
)

func init() {
	flag.Parse()
}

func main() {
	glog.Infof("Start monbus with version %s", version)
	w, err := watcher.New()
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
