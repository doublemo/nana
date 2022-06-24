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

// IsModilePhone 手机号验证
// ^1([358][0-9]|4[579]|66|7[01235678]|9[189])[0-9]{8}$
func IsModilePhone(s string) bool {
	return regexp.MustCompile(`^(1[3-9])\d{9}$`).MatchString(s)
}
