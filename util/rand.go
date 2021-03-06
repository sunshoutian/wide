// Copyright (c) 2014-2018, b3log.org & hacpai.com
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

package util

import (
	"math/rand"
	"time"
)

type myrand struct{}

// Random utilities.
var Rand = myrand{}

// String returns a random string ['a', 'z'] in the specified length
func (*myrand) String(length int) string {
	bytes := make([]byte, length)

	for i := 0; i < length; i++ {
		bytes[i] = byte(Rand.Int('a', 'z'))

		time.Sleep(100 * time.Nanosecond)
	}

	return string(bytes)
}

// Int returns a random integer in range [min, max].
func (*myrand) Int(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	time.Sleep(100 * time.Nanosecond)

	return min + rand.Intn(max-min)
}
