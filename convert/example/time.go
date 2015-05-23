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
package main

import (
	"fmt"
	"github.com/aiwuTech/devKit/convert"
)

func main() {
	time1 := convert.StrTime2Unix("2015-04-09 23:59:59", "2006-01-02 15:04:05")
	time2 := convert.StrTime2Unix("2015-04-09", "2006-01-02")
	fmt.Println(time1)
	fmt.Println(time2)
	fmt.Println(time1 - time2)
}
