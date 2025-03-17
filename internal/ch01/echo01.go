package ch01

import (
	"fmt"
	"os"
)

// Echo01 prints its command-line arguments
func Echo01() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// Echo02 prints its command-line arguments
func Echo02() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// Echo03 prints its command-line arguments
func Echo03() {
	fmt.Println(os.Args[1:])
}
