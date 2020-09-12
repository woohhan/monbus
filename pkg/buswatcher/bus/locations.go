package bus

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
)

const (
	RouteIdFor8109 = "234001236"

	busLocationBaseUri = "http://openapi.gbis.go.kr/ws/rest/buslocationservice"
	serviceKey         = "zx6dk8m%2FvOReysfHNTJf6yn9WEf8GE%2FzSfiWtP8pCiKUpLDPgAN04D1M%2FSW25EogOFU4ICL3InHjsMAHeMyrjA%3D%3D"
)

type response struct {
	XMLName      xml.Name     `xml:"response"`
	ComMsgHeader comMsgHeader `xml:"comMsgHeader"`
	MsgHeader    msgHeader    `xml:"msgHeader"`
	MsgBody      msgBody      `xml:"msgBody"`
}

type comMsgHeader struct {
	ErrMsg string `xml:"errMsg"`
}

type msgHeader struct {
	ResultCode    int    `xml:"resultCode"`
	ResultMessage string `xml:"resultMessage"`
}

type msgBody struct {
	BusLocationList []locationList `xml:"busLocationList"`
}

type locationList struct {
	StationSeq int `xml:"stationSeq"`
}

func GetBusLocations(routeId string) ([]int, error) {
	// HTTP call
	uri := getBusLocationUri(serviceKey, routeId)
	xmlResp, err := resty.New().R().Get(uri)
	if err != nil {
		return nil, err
	}

	// XML 파싱
	var resp response
	decoder := xml.NewDecoder(strings.NewReader(xmlResp.String()))
	if err := decoder.Decode(&resp); err != nil {
		return nil, err
	}
	if resp.MsgHeader.ResultCode != 0 {
		return nil, errors.New(resp.ComMsgHeader.ErrMsg)
	}

	// 결과 정리
	loc := make([]int, 0)
	for _, d := range resp.MsgBody.BusLocationList {
		loc = append(loc, d.StationSeq)
	}
	return loc, nil
}

func getBusLocationUri(serviceKey, routeId string) string {
	return fmt.Sprintf("%s?serviceKey=%s&routeId=%s", busLocationBaseUri, serviceKey, RouteIdFor8109)
}
