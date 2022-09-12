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

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			fmt.Printf("j : %d, jarr : %d | \n", j, len(array)-i-1)
			fmt.Printf("array[j] > : %d, array[j+1] : %d | \n", array[j], array[j+1])
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
			fmt.Printf("array : %v \n", array)
		}
	}
	return array
}
func main() {
	array := []int{11, 14, 3, 8, 18, 17, 43}
	fmt.Println(BubbleSort(array))
}
