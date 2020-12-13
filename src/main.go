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
	// node := NewNode(m)
	// node.BeginningOf(`  	2 * (3 + 1)  `)
	// re := regexp.MustCompile(`\d+`)
	// ss := RegexpEx{re}
	// fmt.Println(ss.StartsWith("123nimoh"))
}
