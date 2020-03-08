// 两数之和
package main

import (
	"fmt"
	"sort"
)

func main() {
	list := []int{1, 2, 3}
	target := 3
	fmt.Println(twoSum(list, target))
}
func twoSum(nums []int, target int) []int {
	var rst []int
	var rstIndex []int
	var n = make([]int, len(nums))
	copy(n, nums)
	sort.Ints(n)
	l := len(n)
OuterLoop:
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			sum := n[i] + n[j]
			if sum == target {
				rst = append(rst, n[i])
				rst = append(rst, n[j])
				break OuterLoop
			} else if sum > target {
				break
			}
		}
	}
	var m = make([]int, len(nums))
	for _, v := range rst {
		for k := 0; k < l; k++ {
			if v == nums[k] && m[k] == 0 {
				rstIndex = append(rstIndex, k)
				m[k] = 1
				break
			}
		}
	}
	sort.Ints(rstIndex)
	return rstIndex
}
