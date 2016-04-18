/*
Copyright 2016 Zbigniew Mandziejewicz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package snowflake

import (
	"fmt"
	"testing"
)

func TestSnowflake(t *testing.T) {
	total := 1000000
	ids := make(map[int64]struct{})

	fmt.Println("test")
	sf, err := New(0)
	if err != nil {
		t.Fatal("Failed to create snowflake generator")
	}

	prev := int64(0)
	for i := 0; i < total; i++ {
		id := <-sf

		if id <= prev {
			t.Errorf("Snowflake is not monotonic: %d <= %d", id, prev)
		}
		prev = id

		_, ok := ids[id]
		if ok {
			t.Errorf("Duplicate snowflake: %d", id)
		}
		ids[id] = struct{}{}
	}
}

func TestNextMillisec(t *testing.T) {
	t1 := timestamp()
	t2 := nextMillisec(t1)

	if t2 <= t1 {
		t.Errorf("Time was not advanced to next millisecond")
	}
}
