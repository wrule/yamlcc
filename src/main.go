package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func readBytes(filePath string) []byte {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic("读取语法定义文件出错")
	}
	return bytes
}

func main() {
	bytes := readBytes("/home/gu/桌面/yamlcc/lang/lua/luanew.yaml")
	m := make(map[interface{}]interface{})
	yaml.Unmarshal(bytes, &m)
	root := Compile(m)
	rst := root.GetDef("number").BeginningOf(`123`)
	fmt.Println(rst)
	rst.Print()
}
