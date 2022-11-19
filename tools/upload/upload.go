package upload

import (
	"encoding/json"
	"errors"
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
	var req Request
	req.List = append(req.List, localPath)

	reqBs, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	// 发送请求
	resp, err := http.Post(url, "application/json", strings.NewReader(string(reqBs)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 解析返回值
	var respData Response
	err = json.Unmarshal(bs, &respData)
	if err != nil {
		return "", err
	}

	// 上传失败
	if !respData.Success {
		return "", ErrPicgoUpload
	}

	// 返回值为空
	if len(respData.Result) == 0 {
		return "", ErrPicgoResultIsEmpty
	}

	return respData.Result[0], nil
}
