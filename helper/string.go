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

package helper

import "unicode/utf8"

// ReplaceCardNo 指定字符前部与后部保留长度,替换中间字符
// 如:589********25
func ReplaceCardNo(str, replaceStr string, s, e int) string {
	count := utf8.RuneCountInString(str)
	if count < s+e {
		return str
	}

	var (
		rep     string
		m       string
		j       int
		counter int
	)

	for i := 0; i < len(str); i++ {
		m = str[j : i+1]
		if !utf8.FullRuneInString(m) {
			continue
		}
		j = i + 1
		counter++
		if counter > s && counter <= count-e {
			rep += replaceStr
		} else {
			rep += m
		}
	}

	return rep
}
