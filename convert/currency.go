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

// 分转换成元
func RMBFen2Yuan(c int64) float64 {
	return Float64Trunc(float64(c)/float64(100), 2)
}

func RMBYuan2Fen(c float64) int64 {
	return int64(Float64Trunc(c, 2)*float64(100))
}