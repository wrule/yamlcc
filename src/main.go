package main

import (
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
	node.GetNode("exps").Run(`
		3 * (1 + 2)
	`)
}
