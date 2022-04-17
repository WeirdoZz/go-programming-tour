package api

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type UploadFile struct {
	// 表单名称
	Name string
	// 文件全路径
	Filepath string
}

// 请求客户端
var httpClient = &http.Client{}

func PostFile(reqUrl string, reqParams map[string]string, files []UploadFile, headers map[string]string) string {
	return post(reqUrl, reqParams, "multipart/form-data", files, headers)
}

func post(reqUrl string, reqParams map[string]string, contentType string, files []UploadFile, headers map[string]string) string {
	requestBody, realContentType := getReader(reqParams, contentType, files)
	httpRequest, _ := http.NewRequest("POST", reqUrl, requestBody)
	// 添加请求头
	httpRequest.Header.Add("Content-Type", realContentType)
	if headers != nil {
		for k, v := range headers {
			httpRequest.Header.Add(k, v)
		}
	}
	// 发送请求
	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	return string(response)
}

func getReader(reqParams map[string]string, contentType string, files []UploadFile) (io.Reader, string) {
	// 如果是json格式的文件
	if strings.Index(contentType, "json") > -1 {
		bytesData, _ := json.Marshal(reqParams)
		return bytes.NewReader(bytesData), contentType
	} else if files != nil {
		// 如果要上传文件
		body := &bytes.Buffer{}
		// 文件写入 body
		writer := multipart.NewWriter(body)
		// 遍历需要上传的文件
		for _, uploadFile := range files {
			file, err := os.Open(uploadFile.Filepath)
			if err != nil {
				panic(err)
			}
			// 创建一个FormFile形式的writer，这样就可以向服务器发送数据了
			part, err := writer.CreateFormFile(uploadFile.Name, filepath.Base(uploadFile.Filepath))
			if err != nil {
				panic(err)
			}
			// 将文件复制到要发送给服务器的writer中
			_, err = io.Copy(part, file)
			file.Close()
		}
		// 其他参数列表写入 body
		for k, v := range reqParams {
			if err := writer.WriteField(k, v); err != nil {
				panic(err)
			}
		}
		if err := writer.Close(); err != nil {
			panic(err)
		}
		// 上传文件需要自己专用的contentType
		return body, writer.FormDataContentType()
	} else {
		urlValues := url.Values{}
		for key, val := range reqParams {
			urlValues.Set(key, val)
		}
		reqBody := urlValues.Encode()
		return strings.NewReader(reqBody), contentType
	}
}

func TestUpload_UploadFile(t *testing.T) {
	files := []UploadFile{
		{Name: "file", Filepath: "C:/Users/Weirdo/Desktop/Snipaste_2022-04-14_19-58-26.png"},
	}
	// url，请求参数，文件，header
	url := "http://localhost:9999/upload/file"
	m := map[string]string{"type": "1"}
	//httptest.NewRequest()
	// 返回string
	var response = PostFile(url, m, files, nil)
	print(response)
}
