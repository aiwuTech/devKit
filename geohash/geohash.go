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
package geohash

import "github.com/gansidui/geohash"

const (
	GeohashPrecision_2500      int = iota + 1 // 2500公里
	GeohashPrecision_630                      // 630公里
	GeohashPrecision_78                       // 78公里
	GeohashPrecision_20                       // 20公里
	GeohashPrecision_2_4                      // 2.4公里
	GeohashPrecision_0_61                     // 610米
	GeohashPrecision_0_01911                  // 20米
	GeohashPrecision_0_00478                  // 5米
	GeohashPrecision_0_0005971                // 60厘米
	GeohashPrecision_0_0001492                // 15厘米
	GeohashPrecision_0_0000186                // 2厘米
)

// geohash精度的设定参考 http://en.wikipedia.org/wiki/Geohash
// geohash length	lat bits	lng bits	lat error	lng error	km error
// 1				2			3			±23			±23			±2500
// 2				5			5			± 2.8		± 5.6		±630
// 3				7			8			± 0.70		± 0.7		±78
// 4				10			10			± 0.087		± 0.18		±20
// 5				12			13			± 0.022		± 0.022		±2.4
// 6				15			15			± 0.0027	± 0.0055	±0.61
// 7				17			18			±0.00068	±0.00068	±0.076
// 8				20			20			±0.000085	±0.00017	±0.019
func GEOHash(addr string, precision int) (hash string, lat, lng float64) {
	loc, err := GetGeoViaAddress(addr)
	if err != nil {
		return
	}

	lat = loc.Result.Location.Lat
	lng = loc.Result.Location.Lng
	hash, _ = geohash.Encode(lat, lng, precision)

	return
}

func GetNeighbors(lat, lng float64, precision int) []string {
	return geohash.GetNeighbors(lat, lng, precision)
}
