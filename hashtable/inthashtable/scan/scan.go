package main

import (
	. "fmt"
	"strings"
)

func main() {
	var input string
	var eof rune = 26

	for {
		_, _ = Scanf("%s\r\n", &input)
		if strings.Contains(input, string(eof)) {
			Printf("输入结束\n")
			break
		}
	}
}
