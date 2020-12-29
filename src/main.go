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
	bytes := readBytes("/home/gu/桌面/yamlcc/lang/lua/luanew.yaml")
	m := make(map[interface{}]interface{})
	yaml.Unmarshal(bytes, &m)
	root := node.Compile(m)
	rst := root.BeginningTrimOf(`  1234`)
	fmt.Printf("match: %s\nnext: %s\nsuccess: %v\n", rst.Match(), rst.Next(), rst.Success())
}
