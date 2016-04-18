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
	"errors"
	"time"
)

const (
	workerIDOffset  = 12
	timestampOffset = 22
	maxTimestamp    = 1<<41 - 1
	maxWorkerID     = 1<<10 - 1
	maxSequence     = 1<<12 - 1
)

var (
	// ErrInvalidWorkerID indicates that worker id is out of valid bounds
	ErrInvalidWorkerID = errors.New("Worker ID is not within bounds")
	since              = time.Date(2012, 1, 0, 0, 0, 0, 0, time.UTC).UnixNano()
)

// SnowFlake is a monotonic ID generator inspired by Twitter Snowflake
// ID is composed of:
//   - unused sign bit
//   - 41 bits of timestamp
//   - 10 bits of worker ID
//   - 12 bits of sequence number
type SnowFlake <-chan int64

// New constructs generator for snowflake IDs
// ErrInvalidWorkerID is returned if WorkerID is not fitting in 10 bits
func New(workerID uint64) (SnowFlake, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, ErrInvalidWorkerID
	}

	sf := make(chan int64)
	go func() {
		last := timestamp()
		seq := uint64(0)
		for {
			ts := timestamp()
			if ts < last {
				ts = nextMillisec(last)
			}

			if ts != last {
				seq = 0
				last = ts
			} else if seq == maxSequence {
				ts = nextMillisec(ts)
				seq = 0
				last = ts
			} else {
				seq++
			}

			id := int64((ts << timestampOffset) | (workerID << workerIDOffset) | seq)
			sf <- id
		}
	}()

	return sf, nil
}

func nextMillisec(ts uint64) uint64 {
	i := timestamp()
	for ; i <= ts; i = timestamp() {
		time.Sleep(100 * time.Microsecond)
	}
	return i
}

func timestamp() uint64 {
	return (uint64(time.Now().UnixNano()-since) / uint64(time.Millisecond)) & maxTimestamp
}
