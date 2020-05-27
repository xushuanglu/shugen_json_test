package controllers

import (
	"fmt"
	"testing"
)

func TestCode(t *testing.T) {

	//正确的写法
	var y interface{} = nil
	_ = y

	x := []string{"a", "b", "c"}
	//错误的写法  0 1 2
	for v := range x {
		fmt.Println(v) // 1 2 3
	}

	//正确的写法 a b c
	for _, v := range x {
		fmt.Println(v) // 1 2 3
	}

	// 错误的 key 检测方式
	a := map[string]string{"one": "2", "two": "", "three": "3"}
	if v := a["two"]; v == "" {
		fmt.Println("key two is no entry") // 键 two 存不存在都会返回的空字符串
	}

	// 正确示例
	if _, ok := a["two"]; !ok {
		fmt.Println("key two is no entry") // 键 two 存不存在都会返回的空字符串
	}

	//错误写法
	// x1 := "text"
	// x1[0] = "T" // error: cannot assign to x[0]
	// fmt.Println(x1)

	//正确写法
	x1 := "text"
	xBytes := []byte(x1)
	fmt.Println("xBytes:", xBytes)
	xBytes[0] = 'T' // 注意此时的 T 是 rune 类型
	x1 = string(xBytes)
	fmt.Println(x1) // Text

	xRunes := []rune(x1)
	xRunes[0] = '我'
	x1 = string(xRunes)
	fmt.Println(x1) // 我ext

}
