// 13. 罗马数字转整数

package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	relation := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var rst int
	var p []string
	for _, i := range s {
		p = append(p, string(i))
	}
	for k, v := range p {
		if k == 0 {
			rst = relation[v]
		} else {
			if relation[v] <= relation[p[k-1]] {
				rst += relation[v]
			} else {
				if p[k-1] == "I" {
					if v == "V" || v == "X" {
						rst = rst - 2*relation[p[k-1]] + relation[v]
					} else {
						return -1
					}
				} else if p[k-1] == "X" {
					if v == "L" || v == "C" {
						rst = rst - 2*relation[p[k-1]] + relation[v]
					} else {
						return -1
					}
				} else if p[k-1] == "C" {
					if v == "D" || v == "M" {
						rst = rst - 2*relation[p[k-1]] + relation[v]
					} else {
						return -1
					}
				} else {
					return -1
				}
			}
		}
	}
	return rst
}
