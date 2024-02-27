// Copyright 2019 The Bitalostored author and other contributors.
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
	"reflect"
	"testing"
)

func TestContainsIndex(t *testing.T) {
	for id, test := range []struct {
		prefix   string
		not      bool
		fixture  string
		index    int
		segments []int
	}{
		{
			"ab",
			false,
			"abc",
			0,
			[]int{2, 3},
		},
		{
			"ab",
			false,
			"fffabfff",
			0,
			[]int{5, 6, 7, 8},
		},
		{
			"ab",
			true,
			"abc",
			0,
			[]int{0},
		},
		{
			"ab",
			true,
			"fffabfff",
			0,
			[]int{0, 1, 2, 3},
		},
	} {
		p := NewContains(test.prefix, test.not)
		index, segments := p.Index(test.fixture)
		if index != test.index {
			t.Errorf("#%d unexpected index: exp: %d, act: %d", id, test.index, index)
		}
		if !reflect.DeepEqual(segments, test.segments) {
			t.Errorf("#%d unexpected segments: exp: %v, act: %v", id, test.segments, segments)
		}
	}
}

func BenchmarkIndexContains(b *testing.B) {
	m := NewContains(string(bench_separators), true)

	for i := 0; i < b.N; i++ {
		_, s := m.Index(bench_pattern)
		releaseSegments(s)
	}
}

func BenchmarkIndexContainsParallel(b *testing.B) {
	m := NewContains(string(bench_separators), true)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, s := m.Index(bench_pattern)
			releaseSegments(s)
		}
	})
}
