// 回文数
// 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
package main

import "fmt"

func main() {
	x := 10
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := x
	var r int
	for {
		if s == 0 {
			break
		}
		r = r*10 + s%10
		s = s / 10
	}
	return r == x
}
