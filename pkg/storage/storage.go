package storage

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"reflect"
	"strings"
	"time"
)

// Storage 는 DB에 정류장 번호와 시간을 저장합니다
type Storage struct {
	db *sql.DB
}

func New() (*Storage, error) {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/mysql?charset=utf8&parseTime=True&loc=Asia%2FSeoul")
	if err != nil {
		return nil, err
	}
	if err := testDbConnection(db); err != nil {
		if !reflect.DeepEqual(err, sql.ErrNoRows) {
			return nil, err
		}
	}
	return &Storage{db: db}, nil
}

func testDbConnection(db *sql.DB) error {
	var id int
	return db.QueryRow("select id from bustime").Scan(&id)
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

// GetBusTime 은 오늘의 버스 시간을 얻어옵니다
func (s *Storage) GetBusTime(stationId string) (string, error) {
	rows, err := s.db.Query("select dateTime from bustime where DATE(datetime) = CURDATE() AND stationId = ? order by datetime", stationId)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	result := make([]string, 0)
	for rows.Next() {
		var t time.Time
		if err := rows.Scan(&t); err != nil {
			return "", err
		}
		result = append(result, t.Format(time.Kitchen))
	}
	return strings.Join(result, "\n"), nil
}
