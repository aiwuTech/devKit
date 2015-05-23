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
	"log"
	"time"
)

const (
	StdDateTimeLayout = "2006-01-02 15:04:05"
	StdDateLayout     = "2006-01-02"
	StdTimeLayout     = "15:04:05"
)

const (
	Second int64 = 1
	Minute       = Second * 60
	Hour         = Minute * 60
	Day          = Hour * 24
)

func StrTime2Unix(strTime string, layout string) int64 {
	t, e := time.Parse(layout, strTime)
	if e != nil {
		log.Printf("<convert> parse time err: %+v\n", e)
		return 0
	}

	return t.Unix()
}
