package main

import (
	"fmt"
	"io/ioutil"
	"regexp"

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
	fmt.Println(node.Keys())

	re := regexp.MustCompile(`^\d+/`)
	fmt.Println(re.MatchString("21/"))
}
