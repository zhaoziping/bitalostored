// Copyright 2019-2024 Xu Ruibo (hustxurb@163.com) and Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package match

import (
	"fmt"
)

type Nothing struct{}

func NewNothing() Nothing {
	return Nothing{}
}

func (self Nothing) Match(s string) bool {
	return len(s) == 0
}

func (self Nothing) Index(s string) (int, []int) {
	return 0, segments0
}

func (self Nothing) Len() int {
	return lenZero
}

func (self Nothing) String() string {
	return fmt.Sprintf("<nothing>")
}
