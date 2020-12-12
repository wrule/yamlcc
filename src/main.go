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
	node.GetDefNode("exps").Test(`  	2 * (3 + 1)  `)
	// re := regexp.MustCompile(`\d+`)
	// fmt.Println(re.SubexpIndex("d1234 sdf12"))
}
