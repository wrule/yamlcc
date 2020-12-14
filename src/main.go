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
	// dict := node.NewDict(m)
	// fmt.Println(dict)
	a := node.NewReg(`\d+`)
	b := node.NewReg(`\s+`)
	a.SetNext(b)
	fmt.Printf("%v\n", a)
}
