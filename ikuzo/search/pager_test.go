// Copyright 2020 Delving B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package search

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewScrollPager(t *testing.T) {
	tests := []struct {
		name string
		want ScrollPager
	}{
		{
			"new pager with defaults",
			ScrollPager{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewScrollPager()
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("NewScrollPager() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
