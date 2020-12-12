package main

import "regexp"

// RegexpEx 自己拓展的正则表达式
type RegexpEx struct {
	*regexp.Regexp
}

// StartsWith 匹配头，并且返回匹配到的字符串和剩余字符串
func (me *RegexpEx) StartsWith(text string) (string, string) {
	indexs := me.FindStringIndex(text)
	if len(indexs) > 1 && indexs[0] == 0 {
		return text[:indexs[1]], text[indexs[1]:]
	}
	return "", text
}
