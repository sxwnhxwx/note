package main

//3-6 字符和字符串处理
/*
	在Golang中string类型的底层是通过byte数组实现的,
	在unicode编码中,中文字符占两个字节,而在utf-8编码中,中文字符占三个字节而Golang的默认编码正是utf-8.
	如果想要获得真实的字符串长度而不是其所占用字节数,有两种方法实现
	1, fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))
	2, fmt.Println("rune:", len([]rune(str)))
*/
import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "abc键盘!"
	for k, v := range []byte(s) { // utf-8 每个字母占1字节，每个汉字占3字节
		fmt.Printf("(%d %X) ", k, v)
	}
	fmt.Println()
	for k, ch := range s { // ch is a rune, unicode
		fmt.Printf("(%d %X) ", k, ch)
	}
	fmt.Println()

	fmt.Println("String count:", len(s))
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	// DecodeRune(p []byte) (r rune, size int)
	// unpacks the first UTF-8 encoding in p and returns the rune and its width in bytes
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		fmt.Printf("(%c %d) ", ch, size)
		bytes = bytes[size:]
	}
	fmt.Println()
	//使用rune
	for i, ch := range []rune(s) { //重新解码为rune数组,rune是int32的别名，因此每个rune占4个字节
		fmt.Printf("(%d %c) ", i, ch)
	}
}

/*
(0 61) (1 62) (2 63) (3 E9) (4 94) (5 AE) (6 E7) (7 9B) (8 98) (9 21)
(0 61) (1 62) (2 63) (3 952E) (6 76D8) (9 21)
String count: 10
Rune count: 6
(a 1) (b 1) (c 1) (键 3) (盘 3) (! 1)
(0 a) (1 b) (2 c) (3 键) (4 盘) (5 !)
*/
