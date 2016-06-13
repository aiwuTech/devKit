// Copyright 2016 zm@huantucorp.com
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
/*
                   _ooOoo_
                  o8888888o
                  88" . "88
                  (| -_- |)
                  O\  =  /O
               ____/`---'\____
             .'  \\|     |//  `.
            /  \\|||  :  |||//  \
           /  _||||| -:- |||||-  \
           |   | \\\  -  /// |   |
           | \_|  ''\---/''  |   |
           \  .-\__  `-`  ___/-. /
         ___`. .'  /--.--\  `. . __
      ."" '<  `.___\_<|>_/___.'  >'"".
     | | :  `- \`.;`\ _ /`;.`/ - ` : | |
     \  \ `-.   \_ __\ /__ _/   .-` /  /
======`-.____`-.___\_____/___.-`____.-'======
                   `=---='
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
         佛祖保佑       永无BUG
*/
package geohash

import "math"

/**
 * 各地图API坐标系统比较与转换;
 * WGS84坐标系：即地球坐标系，国际上通用的坐标系。设备一般包含GPS芯片或者北斗芯片获取的经纬度为WGS84地理坐标系,
 * 谷歌地图采用的是WGS84地理坐标系（中国范围除外）;
 * GCJ02坐标系：即火星坐标系，是由中国国家测绘局制订的地理信息系统的坐标系统。由WGS84坐标系经加密后的坐标系。
 * 谷歌中国地图和搜搜中国地图采用的是GCJ02地理坐标系; BD09坐标系：即百度坐标系，GCJ02坐标系经加密后的坐标系;
 * 搜狗坐标系、图吧坐标系等，估计也是在GCJ02基础上加密而成的。 chenhua
 */
const (
	a  float64 = 6378245.0
	ee float64 = 0.00669342162296594323
)

// 84 to 火星坐标系 (GCJ-02) World Geodetic System ==> Mars Geodetic System
func Wgs842Gcj02(lat, lng float64) *Point {
	if outOfChina(lat, lng) {
		return nil
	}

	return transform(lat, lng)
}

// 火星坐标系 (GCJ-02) to 84 * * @param lon * @param lat * @return
func Gcj022Wgs84(lat, lng float64) *Point {
	point := transform(lat, lng)
	lontitude := lng*2 - point.Lng()
	latitude := lat*2 - point.Lat()
	return NewPoint(latitude, lontitude)
}

// 火星坐标系 (GCJ-02) 与百度坐标系 (BD-09) 的转换算法 将 GCJ-02 坐标转换成 BD-09 坐标
func Gcj022Bd09(lat, lng float64) *Point {
	x := lng
	y := lat
	z := math.Sqrt(x*x+y*y) + 0.00002*math.Sin(y*math.Pi)
	theta := math.Atan2(y, x) + 0.000003*math.Cos(x*math.Pi)
	bd_lng := z*math.Cos(theta) + 0.0065
	bd_lat := z*math.Sin(theta) + 0.006
	return NewPoint(bd_lat, bd_lng)
}

// 火星坐标系 (GCJ-02) 与百度坐标系 (BD-09) 的转换算法 * * 将 BD-09 坐标转换成GCJ-02 坐标
func Bd092Gcj02(lat, lng float64) *Point {
	x := lng - 0.0065
	y := lat - 0.006
	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*math.Pi)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*math.Pi)
	gg_lng := z * math.Cos(theta)
	gg_lat := z * math.Sin(theta)
	return NewPoint(gg_lat, gg_lng)
}

// Bd09 -- wgs84
func Bd092Wgs84(lat, lng float64) *Point {
	gcj02 := Bd092Gcj02(lat, lng)
	map84 := Gcj022Wgs84(gcj02.Lat(), gcj02.Lng())
	return map84
}

func outOfChina(lat, lng float64) bool {
	if lng < 72.004 || lng > 137.8347 {
		return true
	}

	if lat < 0.8293 || lat > 55.8271 {
		return true
	}

	return false
}

func transformLat(x, y float64) float64 {
	ret := -100.0 + 2.0*x + 3.0*y + 0.2*y*y + 0.1*x*y + 0.2*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*math.Pi) + 20.0*math.Sin(2.0*x*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(y*math.Pi) + 40.0*math.Sin(y/3.0*math.Pi)) * 2.0 / 3.0
	ret += (160.0*math.Sin(y/12.0*math.Pi) + 320*math.Sin(y*math.Pi/30.0)) * 2.0 / 3.0
	return ret
}

func transformLng(x, y float64) float64 {
	ret := 300.0 + x + 2.0*y + 0.1*x*x + 0.1*x*y + 0.1*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*math.Pi) + 20.0*math.Sin(2.0*x*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(x*math.Pi) + 40.0*math.Sin(x/3.0*math.Pi)) * 2.0 / 3.0
	ret += (150.0*math.Sin(x/12.0*math.Pi) + 300.0*math.Sin(x/30.0*math.Pi)) * 2.0 / 3.0
	return ret
}

func transform(lat, lng float64) *Point {
	if outOfChina(lat, lng) {
		return NewPoint(lat, lng)
	}
	dLat := transformLat(lng-105.0, lat-35.0)
	dLng := transformLng(lng-105.0, lat-35.0)
	radLat := lat / 180.0 * math.Pi
	magic := math.Sin(radLat)
	magic = 1 - ee*magic*magic
	sqrtMagic := math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * math.Pi)
	dLng = (dLng * 180.0) / (a / sqrtMagic * math.Cos(radLat) * math.Pi)
	mgLat := lat + dLat
	mgLng := lng + dLng
	return NewPoint(mgLat, mgLng)
}
