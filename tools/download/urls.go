package download

import "net/url"

// TrimUrl 删除 URL 中的参数
func TrimUrl(urlStr string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	return u.Scheme + "://" + u.Host + u.Path
}
