package regs

import (
	"regexp"
	"strings"
)

var imgUrlReg = regexp.MustCompile(`!\[.*\]\((.*)\)`)

// GetMarkdownImageUrls 获取markdown中的图片链接
func GetMarkdownImageUrls(text string) []string {
	sss := imgUrlReg.FindAllStringSubmatch(text, -1)

	var urls []string
	for _, ss := range sss {
		for _, s := range ss {
			if strings.HasPrefix(s, "http") {
				urls = append(urls, s)
			}
		}
	}

	return urls
}
