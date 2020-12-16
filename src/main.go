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
		panic("读取文件出错")
	}
	m := make(map[interface{}]interface{})
	yaml.Unmarshal(bytes, &m)
	// fmt.Println(m)
	dict := node.BuildLeafNode(m)
	// matchingText, nextText, success := dict.BeginningTrimOf(`1+(2*3 / (4 / 5 ) + 1991)`)
	matchingText, nextText, success := dict.BeginningTrimOf(`
		if 1 + 2 then
			7897878
		elseif 7 + 8 then
			898989
		elseif 8 then
			18792
		else
			787878
		end
	`)
	fmt.Println("匹配到:", matchingText)
	fmt.Println("剩下的:", nextText)
	fmt.Println("结果:", success)
}
