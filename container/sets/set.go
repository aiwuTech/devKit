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
package sets

type Set interface {
	Add(e interface{}) bool
	Remove(e interface{})
	Clear()
	Contains(e interface{}) bool
	Len() int
	Same(other Set) bool
	Elements() []interface{}
	String() string
}

// 判断集合 one 是否是集合 other 的超集
func IsSuperset(one, other Set) bool {
	if one == nil || other == nil {
		return false
	}

	oneLen := one.Len()
	otherLen := other.Len()
	if oneLen == 0 || oneLen == otherLen {
		return false
	}

	if oneLen > 0 && otherLen == 0 {
		return true
	}

	for _, v := range other.Elements() {
		if !one.Contains(v) {
			return false
		}
	}

	return true
}

// 求集合 one,other的并集
func Union(one, other Set) Set {
	if one == nil || other == nil {
		return nil
	}

	unionedSet := NewSimpleSet()
	for _, v := range one.Elements() {
		unionedSet.Add(v)
	}

	if other.Len() == 0 {
		return unionedSet
	}

	for _, v := range other.Elements() {
		unionedSet.Add(v)
	}

	return unionedSet
}

// 求集合 one，other的交集
func Intersect(one, other Set) Set {
	if one == nil || other == nil {
		return nil
	}

	intersectedSet := NewSimpleSet()
	if other.Len() == 0 {
		return intersectedSet
	}

	if one.Len() < other.Len() {
		for _, v := range one.Elements() {
			if other.Contains(v) {
				intersectedSet.Add(v)
			}
		}
	} else {
		for _, v := range other.Elements() {
			if one.Contains(v) {
				intersectedSet.Add(v)
			}
		}
	}

	return intersectedSet
}

// 求集合 one，other的差集
func Difference(one, other Set) Set {
	if one == nil || other == nil {
		return nil
	}

	differencedSet := NewSimpleSet()
	for _, v := range one.Elements() {
		if !other.Contains(v) {
			differencedSet.Add(v)
		}
	}

	return differencedSet
}

// 求集合 one，other的对称差集
func SymmetricDifference(one, other Set) Set {
	if one == nil || other == nil {
		return nil
	}

	diffA := Difference(one, other)
	diffB := Difference(other, one)

	return Union(diffA, diffB)
}

func NewSimpleSet() Set {
	return NewHashSet()
}

func IsSet(v interface{}) bool {
	if _, ok := v.(Set); ok {
		return true
	}

	return false
}
