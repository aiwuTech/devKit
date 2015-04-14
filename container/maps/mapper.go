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

import "reflect"

type Mapper interface {
	// 获取键值对应的元素值, 没有则返回nil
	Get(key interface {}) interface {}
	// 添加键值对，并返回旧的元素值，若没有则返回nil，true
	Put(key interface {}, elem interface {}) (interface {}, bool)
	// 删除键值对，返回旧的元素值，若没有返回nil
	Remove(key interface {}) interface {}
	// 清除所有键值对
	Clear()
	// 获取键值对的数量
	Len() int
	// 判断是否包含给定的键值
	Contains(key interface {}) bool
	// 获取所有键值
	Keys() []interface {}
	// 获取所有元素值
	Elems() []interface {}
	// 键值对的字典
	ToMap() map[interface {}]interface {}
	// 获取键的类型
	KeyType() reflect.Type
	// 获取元素的类型
	ElemType() reflect.Type
	// 格式化输出
	String() string
}
