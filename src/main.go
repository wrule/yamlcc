package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func main() {
	bytes, err := ioutil.ReadFile("lang/lua/lua.yaml")
	if err != nil {
		panic("读取文件出错")
	}
	m := make(map[interface{}]interface{})
	yaml.Unmarshal(bytes, &m)
	node := NewNode(m)
	node.Run(" (1 + 2) * 3")
	fmt.Println(node.GetNode("$exps").GetNode("$number").GetNode("$opr"))
}
