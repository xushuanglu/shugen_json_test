package controllers

import (
	"fmt"
	"reflect"
	"testing"
)

type data struct {
	num    int
	checks [10]func() bool   // 无法比较
	doIt   func() bool       // 无法比较
	m      map[string]string // 无法比较
	bytes  []byte            // 无法比较
}

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

	//range 迭代 string 得到的值
	// data := "A\xfe\x02\xff\x04"
	// for _, v := range data {
	// 	fmt.Printf("%#x ", v) // 0x41 0xfffd 0x2 0xfffd 0x4    // 错误
	// }

	// for _, v := range []byte(data) {
	// 	fmt.Printf("%#x ", v) // 0x41 0xfe 0x2 0xff 0x4    // 正确
	// }

	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(v1, v2)) // true

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(m1, m2)) // true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	// 注意两个 slice 相等，值和顺序必须一致
	fmt.Println("v1 == v2: ", reflect.DeepEqual(s1, s2)) // true

	type Operation int

	const (
		// Add Operation = iota //错误
		Add Operation = iota + 1 //正确
		Subtract
		Multiply
	)

	fmt.Println("Add:", Add, "Subtract:", Subtract, "Multiply:", Multiply)

	// Add=0, Subtract=1, Multiply=2

}
