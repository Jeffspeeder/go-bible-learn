package ch02

import "fmt"

/*
 * boiling.go
 * 打印水的沸点
 * %d  十进制整数
 * %x, %o, %b  十六进制，八进制，二进制整数
 * %f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
 * %t  布尔：true或false
 * %c  字符（rune） (Unicode码点)
 * %s  字符串
 * %q  带双引号的字符串"abc"或字符'c'
 * %v  变量的自然形式
 * %T  变量的类型
 * %%  字面上的百分号标志
 */
const boilingF = 212.0

func Boiling() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Boiling point = %g°F or %g°C\n", f, c)
}