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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSnowflake(t *testing.T) {
	total := 1000000
	ids := make(map[uint64]struct{})

	fmt.Println("test")
	sf, err := New(0)
	require.NotNil(t, sf)
	require.NoError(t, err)

	prev := uint64(0)
	for i := 0; i < total; i++ {
		id := <-sf

		assert.True(t, id > prev, "Snowflake is not monotonic")
		prev = id

		_, ok := ids[id]
		require.False(t, ok, "Duplicate snowflake")
		ids[id] = struct{}{}
	}
}
