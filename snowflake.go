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
	maxWorkerID     = 1<<10 - 1
	maxSequence     = 1<<12 - 1
)

var (
	// ErrInvalidWorkerID indicates that worker id is out of valid bounds
	ErrInvalidWorkerID = errors.New("Worker ID is not within bounds")
	since              = time.Date(2012, 1, 0, 0, 0, 0, 0, time.UTC).UnixNano() - nanotime()
)

// SnowFlake is a monotonic ID generator inspired by Twitter Snowflake
type SnowFlake <-chan uint64

// New constructs generator for snowflake IDs
func New(workerID uint64) (SnowFlake, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, ErrInvalidWorkerID
	}

	sf := make(chan uint64)
	go func() {
		last := timestamp()
		seq := uint64(0)
		for {
			ts := timestamp()
			if ts < last {
				panic("Time is not monotonic.")
			}

			if ts != last {
				seq = 0
				last = ts
			} else if seq == maxSequence {
				for ts == timestamp() {
					time.Sleep(100 * time.Microsecond)
				}
				seq = 0
				ts = timestamp()
				last = ts
			} else {
				seq++
			}

			id := (ts << timestampOffset) | (workerID << workerIDOffset) | seq
			sf <- id
		}
	}()

	return sf, nil
}

func timestamp() uint64 {
	return uint64((nanotime() - since) / int64(time.Millisecond))
}
