package watcher

import (
	"github.com/golang/glog"
	"monbus/pkg/bus"
	"monbus/pkg/storage"
	"time"
)

type station struct {
	id            int
	Name          string
	nextWatchTime *time.Time
}

type StationWatcher struct {
	stations map[int]*station
	storage  *storage.Storage
}

func New() (*StationWatcher, error) {
	s, err := storage.New()
	if err != nil {
		return nil, err
	}
	return &StationWatcher{
		stations: map[int]*station{},
		storage:  s,
	}, nil
}

func (s *StationWatcher) Close() error {
	return s.storage.Close()
}

// AddStation 는 감시할 정류장을 추가합니다
func (s *StationWatcher) AddStation(id int, name string) {
	s.stations[id] = &station{
		id:            id,
		Name:          name,
		nextWatchTime: nil,
	}
}

// Watch 는 주어진 정류장을 watchInterval 주기마다 감시합니다. 만약 정류장을 찾으면 Storage에 기록하고 ignoreTime 동안은 무시합니다
func (s *StationWatcher) Watch(watchInterval time.Duration, ignoreTime time.Duration) {
	for {
		locs, err := bus.GetBusLocations(bus.RouteIdFor8109)
		if err != nil {
			glog.Error(err)
		}
		glog.V(2).Infof("GetBusLocations result: %v", locs)
		for _, loc := range locs {
			station, found := s.stations[loc]
			// 기록할 대상이 아니거나 아직 nextWatchTime를 지나지 않았다면 기록하지 않는다
			if !found || (station.nextWatchTime != nil && time.Now().Before(*station.nextWatchTime)) {
				continue
			}
			nextWatchTime := time.Now().Add(ignoreTime)
			station.nextWatchTime = &nextWatchTime
			s.storage.Write(station.id)
		}
		time.Sleep(watchInterval)
	}
}
