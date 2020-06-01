package storage

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"time"
)

// Storage 는 DB에 정류장 번호와 시간을 저장합니다
type Storage struct {
	db *sql.DB
}

func New() (*Storage, error) {
	db, err := sql.Open("mysql", "admin:hRJEgsEbUy94QJmvfjtb@tcp(database-1.c4hts3jaq4u1.ap-northeast-2.rds.amazonaws.com:3306)/mon")
	if err != nil {
		return nil, err
	}
	return &Storage{
		db,
	}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

// Write 는 stationId와 현재 시간을 테이블에 기록합니다
func (s *Storage) Write(stationId int, nowTime time.Time) error {
	glog.V(2).Infof("Writed stationId(%v) dateTime(%v)", stationId, nowTime)
	res, err := s.db.Exec("INSERT INTO bustime (stationId, dateTime) value (?, ?)", stationId, nowTime)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected != 1 {
		return errors.New("write affected is not 1")
	}
	return nil
}
