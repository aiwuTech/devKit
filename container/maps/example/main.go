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
	"github.com/aiwuTech/devKit/container/maps"
	"reflect"
)

func IntCompare(e1, e2 interface{}) int8 {
	k1 := e1.(int)
	k2 := e2.(int)

	if k1 < k2 {
		return -1
	} else if k1 > k2 {
		return 1
	} else {
		return 0
	}
}

func main() {
	keys := maps.NewKeys(IntCompare, reflect.TypeOf(1))
	omap := maps.NewOrderMap(keys, reflect.TypeOf("xxx"))

	omap.Put(166, "xxx")
	omap.Put(4, "sss")
	omap.Put(867, "xxx")
	omap.Put(7, "sdfsdf")

	fmt.Println(omap.Len())
	fmt.Println(omap.Keys())
	fmt.Println(omap.Elems())
	fmt.Println(omap.ToMap())
	fmt.Println(omap.Remove(4))
	fmt.Println(omap.String())
	fmt.Println(omap.Get(166))
	fmt.Println(omap.Remove(867))
}
