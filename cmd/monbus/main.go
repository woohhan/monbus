package main

import (
	"flag"
	"monbus/pkg/watcher"
	"time"

	// "github.com/go-resty/resty/v2"
)

func init() {
	flag.Parse()
}

func main() {
	w, err := watcher.New()
	if err != nil {
		panic(err)
	}
	defer func() {
		w.Close()
	}()
	w.AddStation(4, "미금역")
	w.AddStation(45, "주엽역(중)")
	w.Watch(10*time.Second, 5*time.Minute)
}
