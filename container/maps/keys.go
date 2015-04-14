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
package maps

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
)

// compareFunc的结果值：
//   小于0: 第一个参数小于第二个参数
//   等于0: 第一个参数等于第二个参数
//   大于1: 第一个参数大于第二个参数
type CompareFunction func(interface{}, interface{}) int8

type Keys interface {
	sort.Interface
	Add(elem interface{}) bool
	Remove(elem interface{}) bool
	Clear()
	Get(index int) interface{}
	GetAll() []interface{}
	Search(elem interface{}) (index int, contains bool)
	ElemType() reflect.Type
	CompareFunc() CompareFunction
}

type keys struct {
	container   []interface{}
	compareFunc CompareFunction
	elemType    reflect.Type
}

func (k *keys) Len() int {
	return len(k.container)
}

func (k *keys) Less(i, j int) bool {
	return k.compareFunc(k.container[i], k.container[j]) < 0
}

func (k *keys) Swap(i, j int) {
	k.container[i], k.container[j] = k.container[j], k.container[i]
}

func (k *keys) isAcceptableElem(elem interface{}) bool {
	if elem == nil {
		return false
	}

	if reflect.TypeOf(elem) != k.elemType {
		return false
	}

	return true
}

func (k *keys) Add(elem interface{}) bool {
	if !k.isAcceptableElem(elem) {
		return false
	}

	k.container = append(k.container, elem)
	sort.Sort(k)

	return true
}

func (k *keys) Remove(elem interface{}) bool {
	index, contains := k.Search(elem)
	fmt.Println(index)
	if !contains {
		return false
	}

	k.container = append(k.container[0:index], k.container[index+1:]...)

	return true
}

func (k *keys) Clear() {
	k.container = make([]interface{}, 0)
}

func (k *keys) Get(index int) interface{} {
	if index >= k.Len() {
		return nil
	}

	return k.container[index]
}

func (k *keys) GetAll() []interface{} {
	snapshot := make([]interface{}, 0)
	for _, val := range k.container {
		snapshot = append(snapshot, val)
	}

	return snapshot
}

func (k *keys) Search(elem interface{}) (index int, contains bool) {
	if !k.isAcceptableElem(elem) {
		return
	}

	index = sort.Search(k.Len(), func(i int) bool {
		return k.compareFunc(k.container[i], elem) >= 0
	})

	if index < k.Len() && k.container[index] == elem {
		contains = true
	}

	return
}

func (k *keys) ElemType() reflect.Type {
	return k.elemType
}

func (k *keys) CompareFunc() CompareFunction {
	return k.compareFunc
}

func (k *keys) String() string {
	var buf bytes.Buffer
	buf.WriteString("Keys<")
	buf.WriteString(k.elemType.Kind().String())
	buf.WriteString(">{")
	first := true
	buf.WriteString("[")
	for _, val := range k.container {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", val))
	}
	buf.WriteString("]")
	buf.WriteString("}")

	return buf.String()
}

func NewKeys(compareFunc CompareFunction, elemType reflect.Type) Keys {
	return &keys{
		container:   make([]interface{}, 0),
		compareFunc: compareFunc,
		elemType:    elemType,
	}
}
