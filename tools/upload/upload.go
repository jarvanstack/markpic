package upload

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	url = "http://localhost:36677/upload"
)

var (
	ErrPicgoUpload        = errors.New("picgo 上传失败")
	ErrPicgoResultIsEmpty = errors.New("picgo 返回值为空")
)

type Uploader interface {
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
	return upload(localPath)
}

func upload(localPath string) (string, error) {

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

	return resp.Result[0], nil
}
