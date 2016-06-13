package geohash

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aiwuTech/httpclient"
	"io/ioutil"
	"log"
)

// 百度地图API GEO ： http://developer.baidu.com/map/webservice-geocoding.htm
// 百度地图API IP ： http://developer.baidu.com/map/ip-location-api.htm
// 百度地图API GEO App Key 申请地址：http://lbsyun.baidu.com/apiconsole/key?application=key

// 注意，请到百度地图开发者中心申请自己的 App Key，下方的 key 是百度提供的，不保证永远有效。

var (
	AppKey string = "xwCIsZzaLkFYLkCCfuAKUbkG" // baidu's
	Debug  bool   = true
	_      log.Logger
)

const (
	reqURLForGEO     string = "http://api.map.baidu.com/geocoder/v2/?ak="
	reqURLForIP      string = "http://api.map.baidu.com/location/ip?ak="
	reqURLForGEOCONV string = "http://api.map.baidu.com/geoconv/v1/?ak="
)

const (
	Coords_wgs84  = "wgs84ll" // GPS设备获取的角度坐标，wgs84坐标;
	Coords_gcj02  = "gcj02ll"
	Coords_bd09ll = "bd09ll"
	Coords_bd09mc = "bd09mc"
)

//
//func ConvGEO(lat, lng float64, from, to CoordsType) (*StructGEOCONV, error) {
//	res := new(StructGEOCONV)
//
//	parameter := fmt.Sprintf("&coords=%v,%v&from=%v&to=%v&output=json", lat, lng, from, to)
//	reqURL := fmt.Sprintf("%s%s%s", reqURLForGEOCONV, AppKey, parameter)
//
//	if Debug {
//		log.Printf("baidu request url: %v", reqURL)
//	}
//
//	res2, err := requestBaidu("GEOCONV", reqURL)
//	if Debug {
//		log.Printf("baidu geo conv resp: %v, err: %v", res2, err)
//	}
//	if err != nil {
//		return res, err
//	}
//
//	if res2.(*StructGEOCONV).Status != 0 {
//		message := fmt.Sprintf("百度 API 报错：status: %v, msg:%v", res2.(*StructGEOCONV).Status, res2.(*StructGEOCONV).Msg)
//		return res, errors.New(message)
//	}
//
//	res3 := res2.(*StructGEOCONV)
//	return res3, nil
//}

func GetAddressViaGEO(lat, lng, coordtype string) (*StructGEOToAddress, error) {
	res := new(StructGEOToAddress)

	parameter := fmt.Sprintf("&location=%s,%s&output=json&pois=0&coordtype=%v", lat, lng, coordtype)
	reqURL := fmt.Sprintf("%s%s%s", reqURLForGEO, AppKey, parameter)

	res2, err := requestBaidu("GetAddressViaGEO", reqURL)
	if err != nil {
		return res, err
	}

	if res2.(*StructGEOToAddress).Status != 0 {
		message := fmt.Sprintf("百度 API 报错：%s", res2.(*StructGEOToAddress).Msg)
		return res, errors.New(message)
	}

	res3 := res2.(*StructGEOToAddress)
	return res3, nil
}

func GetGeoViaAddress(address string) (*StructAddressToGEO, error) {
	res := new(StructAddressToGEO)

	parameter := fmt.Sprintf("&address=%s&output=json&pois=1", address)
	reqURL := fmt.Sprintf("%s%s%s", reqURLForGEO, AppKey, parameter)

	res2, err := requestBaidu("GetGeoViaAddress", reqURL)
	if err != nil {
		return res, err
	}

	if res2.(*StructAddressToGEO).Status != 0 {
		message := fmt.Sprintf("百度 API 报错：%s", res2.(*StructAddressToGEO).Msg)
		return res, errors.New(message)
	}

	res3 := res2.(*StructAddressToGEO)
	return res3, nil
}

func GetAddressViaIP(address string) (*StructIPToAddress, error) {
	res := new(StructIPToAddress)

	parameter := fmt.Sprintf("&ip=%s&output=json&pois=0", address)
	reqURL := fmt.Sprintf("%s%s%s", reqURLForIP, AppKey, parameter)

	res2, err := requestBaidu("GetAddressViaIP", reqURL)
	if err != nil {
		return res, err
	}

	if res2.(*StructIPToAddress).Status != 0 {
		message := fmt.Sprintf("百度 API 报错：%s", res2.(*StructIPToAddress).Message)
		return res, errors.New(message)
	}

	res3 := res2.(*StructIPToAddress)
	return res3, nil

}

func requestBaidu(reqType, reqURL string) (interface{}, error) {

	res, err := getResStruct(reqType)
	if err != nil {
		return res, err
	}

	resp, err := httpclient.ForwardHttp("GET", reqURL, nil)

	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 200 {

		err := json.Unmarshal(bytes, &res)

		if err != nil {
			return res, err
		}

	} else {
		return res, errors.New("请求百度API失败，状态码不等于200")
	}

	return res, nil
}

func getResStruct(reqType string) (interface{}, error) {
	var res interface{}

	if reqType == "GetAddressViaIP" {
		return new(StructIPToAddress), nil
	}

	if reqType == "GetGeoViaAddress" {
		return new(StructAddressToGEO), nil
	}

	if reqType == "GetAddressViaGEO" {
		return new(StructGEOToAddress), nil
	}

	if reqType == "GEOCONV" {
		return new(StructGEOCONV), nil
	}
	return res, errors.New("结构体请求错误")
}
