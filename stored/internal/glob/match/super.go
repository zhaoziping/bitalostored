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

type Super struct{}

func NewSuper() Super {
	return Super{}
}

func (self Super) Match(s string) bool {
	return true
}

func (self Super) Len() int {
	return lenNo
}

func (self Super) Index(s string) (int, []int) {
	segments := acquireSegments(len(s) + 1)
	for i := range s {
		segments = append(segments, i)
	}
	segments = append(segments, len(s))

	return 0, segments
}

func (self Super) String() string {
	return fmt.Sprintf("<super>")
}
