package regs

import (
	"regexp"
	"strings"
)

var imgUrlReg = regexp.MustCompile(`!\[.*\]\((.*)\)`)

// GetRemoteImg 获取markdown中的远程图片链接
func GetRemoteImg(text string) []string {
	sss := imgUrlReg.FindAllStringSubmatch(text, -1)

	var urls []string
	for _, ss := range sss {
		for _, s := range ss {
			if strings.HasPrefix(s, "![") {
				continue
			}
			if strings.HasPrefix(s, "http") {
				urls = append(urls, s)
			}
		}
	}

	return urls
}

// GetRemoteImg 获取markdown中的所有本地图片链接
func GetLocalImg(text string) []string {
	sss := imgUrlReg.FindAllStringSubmatch(text, -1)

	var urls []string
	for _, ss := range sss {
		for _, s := range ss {
			if strings.HasPrefix(s, "![") {
				continue
			}
			if strings.HasPrefix(s, "http") {
				continue
			}
			urls = append(urls, s)
		}
	}

	return urls
}
