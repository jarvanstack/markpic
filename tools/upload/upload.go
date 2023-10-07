package upload

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	url = "http://localhost:36677/upload"
)

var (
	ErrPicgoUpload        = errors.New("picgo 上传失败")
	ErrPicgoResultIsEmpty = errors.New("picgo 返回值为空")
)

var (
	ostype = runtime.GOOS
)

type Uploader interface {
	// 路径处理

	// 上传, 返回远程绝对路径
	Upload(localPath string) (string, error)
}

type UploaderImpl struct {
}

type Request struct {
	List []string
}

type Response struct {
	Success bool
	Result  []string
}

func NewUploader() Uploader {
	return &UploaderImpl{}
}

func (u *UploaderImpl) Upload(localPath string) (string, error) {
	// 绝对路径
	localPath, err := filepath.Abs(localPath)
	if err != nil {
		return "", err
	}

	// 统一 window 和 linux 的路径分隔符
	if ostype == "windows" {
		localPath = strings.ReplaceAll(localPath, `/`, `\\`)
		if !strings.HasPrefix(localPath, `\\`) {
			localPath = strings.ReplaceAll(localPath, `\`, `\\`)
		}
	} else {
		localPath = strings.ReplaceAll(localPath, `\\`, "/")
	}

	fmt.Println("[上传] 绝对路径: ", localPath)

	return upload(localPath)
}

// 获取 ipv4 地址
func LocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	var ip string = "localhost"
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	return ip
}

func upload(localPath string) (string, error) {

	// url := "http://" + LocalIp() + ":36677/upload"
	url := "http://127.0.0.1:36677/upload"
	method := "POST"

	data := fmt.Sprintf(`{"list":["%s"]}`, localPath)
	payload := strings.NewReader(data)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var resp Response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if len(resp.Result) == 0 {
		return "", ErrPicgoResultIsEmpty
	}

	return resp.Result[0], nil
}
