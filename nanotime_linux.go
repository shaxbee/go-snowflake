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

/*
#include <time.h>
*/
import "C"

func nanotime() int64 {
	ts := C.struct_timespec{}

	C.clock_gettime(C.CLOCK_MONOTONIC_RAW, &ts)
	return int64(ts.tv_nsec)
}
