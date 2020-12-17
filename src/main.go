package main

import (
	"fmt"
	"io/ioutil"

	"./node"
	"gopkg.in/yaml.v2"
)

func main() {
	bytes, err := ioutil.ReadFile("lang/lua/lua.yaml")
	if err != nil {
		panic("读取语法定义文件出错")
	}
	m := make(map[interface{}]interface{})
	yaml.Unmarshal(bytes, &m)
	dict := node.BuildLeafNode(m)
	matchingText, nextText, success := dict.BeginningTrimOf(`
		3 * (1 + 1992) <= (1 + 2 + 4) * 5 / 123
	`)
	fmt.Println("匹配到:", matchingText)
	fmt.Println("剩下的:", nextText)
	fmt.Println("结果:", success)
}
