// Copyright 2022 lukemakhanu
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

package main

import (
	"fmt"
	"log"
)

func binarySearch(array []int, key int) int {

	max := len(array) - 1
	min := 0

	var mid int
	for min <= max {
		mid = (max + min) / 2
		if array[mid] == key {
			return mid
		} else if array[mid] > key {
			max = mid
		} else {
			min = mid + 1
		}

	}

	return 1
}

func main() {
	fmt.Println("Binary search")
	data := []int{10, 11, 12, 13, 14, 15, 16}
	// do a binary search.
	key := 14
	res := binarySearch(data, key)
	log.Printf("result: %d", res)
}
