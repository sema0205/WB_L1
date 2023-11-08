package main

import (
	"fmt"
	"strings"
)

var justString string

//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}

func someFunc() {
	var builder strings.Builder
	v := "fbcfgaabcacdefgaagaabcfgaabcbcfgaabcacdefgaafgaabcabcfgaabcacdefgaababcdefgaabcdefgfgaabcfgaabcfgaabcfgaabcacdefgaabcdefgaabcdefgafgaabcfgaabcfgaabcababcdefgaabcdefgfgaabcfgaabcfgaabcfgaabcacdefgaabcdefgaabcdefgafgaabcfgaabcfgaabcababcdefgaabcdefgfgaabcfgaabcfgaabcfgaabcacdefgaabcdefgaabcdefga"
	vBytes := []byte(v)
	for i := 0; i < 100; i++ {
		builder.WriteByte(vBytes[i])
	}
	justString = builder.String()
}

func main() {
	someFunc()
	fmt.Println(justString)
}
