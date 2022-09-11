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

import "fmt"

func Selection_Sort(array []int, size int) []int {
	var min_index int
	var temp int
	for i := 0; i < size-1; i++ {
		min_index = i
		// Find index of minimum element
		for j := i + 1; j < size; j++ {
			if array[j] < array[min_index] {
				min_index = j
			}
		}
		temp = array[i]
		fmt.Println("store the current value at that index: ", temp)
		array[i] = array[min_index]
		fmt.Println("assign that index with the least value: ", array)
		array[min_index] = temp
		fmt.Println("finish swapping by assigning the temporary value from where the min value was:", array)
	}
	return array
}
func main() {
	var num = 7
	array := []int{2, 4, 3, 1, 6, 8, 5}
	fmt.Println(Selection_Sort(array, num))
}
