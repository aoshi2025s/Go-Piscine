package main

import (
	"os"
	"fmt"
	"piscine"
)

func main() {
	test()

	args := os.Args
	params := args[1:]
	if piscine.IsValidInput(params) == false {
		return
	}

	value1, err1 := piscine.Atoi(params[0])
	value2, err2 := piscine.Atoi(params[2])
	op := params[1]

	// overflow
	if err1 == false || err2 == false {
		return
	}

	if value2 == 0 && op[0] == '/' {
		fmt.Println("No division by 0")
		return
	} else if value2 == 0 && op[0] == '%' {
		fmt.Println("No modulo by 0")
		return
	}
	
	result, err := piscine.CalcuDoop(value1, value2, op)
	if err == false {
		return
	}
	fmt.Println(result)
}

func test() {
	var params []string
	success := "\033[32m"  // 緑色の開始
	failure := "\033[31m"  // 赤色の開始
	reset := "\033[0m"     // 色をリセット

	if !piscine.IsValidInput(params) {
		fmt.Println(success + "no input test ok" + reset)
	} else {
		fmt.Println(failure + "no input test failed" + reset)
	}

	params = []string{"hello", "+", "1"}
	if !piscine.IsValidInput(params) {
		fmt.Println(success + "not valid input test ok" + reset)
	} else {
		fmt.Println(failure + "not valid input test failed" + reset)
	}

	params = []string{"1", "p", "1"}
	if !piscine.IsValidInput(params) {
		fmt.Println(success + "not valid input test ok" + reset)
	} else {
		fmt.Println(failure + "not valid input test failed" + reset)
	}

	result, err := piscine.CalcuDoop(1, 0, "/")
	if !err {
		fmt.Println(success + "division 0 test ok" + reset)
	} else {
		fmt.Println(failure + "division 0 test failed" + reset)
	}

	result, err = piscine.CalcuDoop(1, 0, "%")
	if !err {
		fmt.Println(success + "modulo 0 test ok" + reset)
	} else {
		fmt.Println(failure + "modulo 0 test failed" + reset)
	}

	result, err = piscine.Atoi("9223372036854775807")
	if err && result == 9223372036854775807 {
		fmt.Println(success + "Atoi overflow test ok" + reset)
	} else {
		fmt.Println(failure + "Atoi overflow test failed" + reset)
	}

	result, err = piscine.Atoi("9223372036854775808")
	if !err {
		fmt.Println(success + "Atoi overflow test ok" + reset)
	} else {
		fmt.Println(failure + "Atoi overflow test failed" + reset)
	}

	result, err = piscine.Atoi("-9223372036854775808")
	if err && result == -9223372036854775808 {
		fmt.Println(success + "Atoi overflow test ok" + reset)
	} else {
		fmt.Println(failure + "Atoi overflow test failed" + reset)
	}

	result, err = piscine.CalcuDoop(9223372036854775807, 1, "+")
	if !err {
		fmt.Println(success + "+ operator overflow test ok" + reset)
	} else {
		fmt.Println(failure + "+ operator overflow test failed" + reset)
	}
	
	result, err = piscine.CalcuDoop(9223372036854775807, -2, "-")
	if !err {
		fmt.Println(success + "- operator overflow test ok" + reset)
	} else {
		fmt.Println(failure + "- operator overflow test failed" + reset)
	}

	result, err = piscine.CalcuDoop(-9223372036854775807, 3, "-")
	if !err {
		fmt.Println(success + "- operator overflow test ok" + reset)
	} else {
		fmt.Println(failure + "- operator overflow test failed" + reset)
	}

	result, err = piscine.CalcuDoop(9223372036854775807, 3, "*")
	if !err {
		fmt.Println(success + "* operator overflow test ok" + reset)
	} else {
		fmt.Println(failure + "* operator overflow test failed" + reset)
	}

	result, err = piscine.CalcuDoop(-9223372036854775808, -1, "/")
	if !err {
		fmt.Println(success + "/ operator overflow test ok" + reset)
	} else {
		fmt.Println(failure + "/ operator overflow test failed" + reset)
	}

	// basic test
	result, err = piscine.CalcuDoop(1, 1, "+")
	if err && result == 2 {
		fmt.Println(success + "basic test1 ok" + reset)
	} else {
		fmt.Println(failure + "basic test1 failed" + reset)
	}

	result, err = piscine.CalcuDoop(1, 1, "*")
	if err && result == 1 {
		fmt.Println(success + "basic test2 ok" + reset)
	} else {
		fmt.Println(failure + "basic test2 failed" + reset)
	}

	result, err = piscine.CalcuDoop(1, -1, "*")
	if err && result == -1 {
		fmt.Println(success + "basic test3 ok" + reset)
	} else {
		fmt.Println(failure + "basic test3 failed" + reset)
	}
}
