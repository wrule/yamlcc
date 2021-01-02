package main

import "fmt"

func main() {
	fmt.Println("你好，世界")
	reNode := NewReg(`\d+`)
	rst := reNode.BeginningOf(` 1234`)
	rst.Print()
}
