// 整数反转
// 给出一个32位的有符号整数，你需要将这个整数中每位上的数字进行反转。
package main

import "fmt"

func main() {
	x := -2147483648
	rst := reverse(x)
	fmt.Println(rst)
}

func reverse(x int) int {
	p := x
	var rst int
	for {
		if rst > 2147483647 || rst < -2147483648 {
			return 0
		}
		if p == 0 {
			break
		}
		rst = rst*10 + p%10
		p = p / 10
	}
	return rst
}
