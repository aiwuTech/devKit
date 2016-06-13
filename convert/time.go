// Copyright 2015 mint.zhao.chiu@gmail.com
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
package convert

import (
	"github.com/jinzhu/now"
	"log"
	"time"
)

const (
	StdDateTimeLayout      = "2006-01-02 15:04:05"
	StdDateLayout          = "2006-01-02"
	StdMM_ddLayout         = "01-02"
	StdLocalDateLayout     = "2006年01月02日"
	StdLocalDateTimeLayout = "2006年01月02日 15时04分05秒"
	StdLocalMonthLayout    = "2006年01月"
	StdLocalMonthDayLayout = "01月02日"
	StdTimeLayout          = "15:04:05"
)

const (
	Second int64 = 1
	Minute       = Second * 60
	Hour         = Minute * 60
	Day          = Hour * 24
	Week         = Day * 7
)

func StrTime2Unix(strTime string, layout string) int64 {
	return Str2Time(strTime, layout).Unix()
}

func Str2Time(strTime, layout string) time.Time {
	if strTime == "" || layout == "" {
		return time.Now()
	}
	t, e := time.ParseInLocation(layout, strTime, time.Local)
	//	t, e := time.Parse(layout, strTime)
	if e != nil {
		log.Printf("<convert> parse time err: %+v\n", e)
		return time.Now()
	}

	return t
}

func StrUnix2Time(unix string) time.Time {
	unixTimestamp := Str2Int64(unix)

	return time.Unix(unixTimestamp, 0)
}

// 是否是今日
func IsToday(timeUnix int64) bool {
	inTime := time.Unix(timeUnix, 0)
	return inTime.After(now.BeginningOfDay()) && inTime.Before(now.EndOfDay())
}

// 是否是昨日
func IsYesterday(timeUnix int64) bool {
	return IsToday(timeUnix + 86400)
}

// 是否是未来
func IsFuture(timeUnix int64) bool {
	return time.Unix(timeUnix, 0).After(time.Now())
}

// 是否是过去
func IsPast(timeUnix int64) bool {
	return !IsFuture(timeUnix)
}

type Cost struct {
	t1 time.Time
	t2 time.Time
}

func NewTimeCost() *Cost {
	return &Cost{}
}

func (c *Cost) Begin() *Cost {
	c.t1 = time.Now()
	return c
}

func (c *Cost) Cost() time.Duration {
	c.t2 = time.Now()
	return c.t2.Sub(c.t1)
}
