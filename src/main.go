package main

import (
	"fmt"
	"io/ioutil"

	"./node"
	"gopkg.in/yaml.v2"
)

func readBytes(filePath string) []byte {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic("读取语法定义文件出错")
	}
	return bytes
}

func readString(filePath string) string {
	return string(readBytes(filePath))
}

func main() {
	bytes := readBytes("/home/gu/桌面/yamlcc/lang/lua/lua.yaml")
	m := make(map[interface{}]interface{})
	yaml.Unmarshal(bytes, &m)
	root := node.Compile(m)
	fmt.Println(root.Nexts()[0].Nexts())
	// for _, node := range nodes {
	// 	node.Print()
	// }
	// code := readString("test/lua/1.lua")
	// matchingText, nextText, success := dict.BeginningTrimOf(code)
	// fmt.Println("匹配到:", matchingText)
	// fmt.Println("剩下的:", nextText)
	// fmt.Println("结果:", success)
	// fmt.Println("你好")
}
