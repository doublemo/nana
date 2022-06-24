// Copyright (c) 2022 The Linna Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validator

import (
	"regexp"
	"strconv"
	"time"
)

// 银行卡校验码
var bankCardNoMask = [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

// IsBankCardNo 银行卡号验证
func IsBankCardNo(s string) bool {
	odd := len(s) & 1
	var sum int
	for i, c := range s {
		if c < '0' || c > '9' {
			return false
		}

		if i&1 == odd {
			sum += bankCardNoMask[c-'0']
		} else {
			sum += int(c - '0')
		}
	}
	return sum%10 == 0
}

// 身证验证
var wi [17]int = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var wix [11]byte = [...]byte{1, 0, 'X', 9, 8, 7, 6, 5, 4, 3, 2}

// IsIDCardNo 身份证号并返回年龄
func IsIDCardNo(id string) (int, bool) {
	pattern := regexp.MustCompile(`^[1-9]\d{5}(18|19|20|21|22|23|24|25)(\d{2})((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
	if !pattern.MatchString(id) {
		return 0, false
	}

	// 强认证
	// 将身份中的前17位变成数字
	idNumArray := make([]int, 17)
	var endfix byte
	for k, v := range []byte(id) {
		if k >= 17 {
			if v == 88 {
				endfix = 88
				continue
			}

			if m, err := strconv.Atoi(string(v)); err == nil {
				endfix = byte(m)
			}
			continue
		}

		idNumArray[k], _ = strconv.Atoi(string(v))
	}

	var res int
	for i := 0; i < 17; i++ {
		res += idNumArray[i] * wi[i]
	}

	code := res % 11
	if code >= len(wix) {
		return 0, false
	}

	if wix[code] != byte(endfix) {
		return 0, false
	}

	matchs := pattern.FindStringSubmatch(id)
	year, err := strconv.Atoi(matchs[1] + matchs[2])
	if err != nil {
		return 0, false
	}

	month, err := strconv.Atoi(matchs[3])
	if err != nil {
		return 0, false
	}

	age := time.Now().Year() - year
	if int(time.Now().Month()) < month {
		age--
	}

	return age, true
}
