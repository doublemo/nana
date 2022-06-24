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

import "regexp"

// IsZip 邮政编码验证
func IsZip(s string) bool {
	re := regexp.MustCompile(`^\d{6}$`)
	return re.MatchString(s)
}

// IsAddress 地址验证
func IsAddress(s string) bool {
	pattern := regexp.MustCompile(`^[a-zA-Z\p{Han}]{1}[a-z0-9A-Z\s\p{Han}]+$`)
	return pattern.MatchString(s)
}
