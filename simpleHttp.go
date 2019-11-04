package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const Sample_GET_URL = "testurl"
const Sample_POST_URL = "testurl"

func main() {
	var (
		err    error
		bs     []byte
		params url.Values
	)
	params = url.Values{}
	params.Set("user_id", strconv.FormatInt(2123, 10))
	params.Set("new_captcha", "2")

	if bs, err = Get(context.Background(), Sample_GET_URL, params); err != nil {
		return
	}
	fmt.Println(bs)
	// post
	if bs, err = Post(context.Background(), Sample_POST_URL, params); err != nil {
		return
	}
	fmt.Println(bs)
}

// NewRequest new a http request.
func NewRequest(method, uri string, params url.Values) (req *http.Request, err error) {
	if method == "GET" {
		req, err = http.NewRequest(method, uri+"?"+params.Encode(), nil)
	} else {
		req, err = http.NewRequest(method, uri, strings.NewReader(params.Encode()))
	}
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return
}

// Do handler request
func Do(c context.Context, req *http.Request) (body []byte, err error) {
	var res *http.Response
	dialer := &net.Dialer{
		Timeout:   time.Duration(1 * int64(time.Second)),
		KeepAlive: time.Duration(1 * int64(time.Second)),
	}
	transport := &http.Transport{
		DialContext:     dialer.DialContext,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: transport,
	}
	req = req.WithContext(c)
	if res, err = client.Do(req); err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode >= http.StatusInternalServerError {
		err = errors.New("http status code 5xx")
		return
	}
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return
	}
	return
}

// Get client.Get send GET request
func Get(c context.Context, uri string, params url.Values) (body []byte, err error) {
	req, err := NewRequest("GET", uri, params)
	if err != nil {
		return
	}
	body, err = Do(c, req)
	return
}

// Post client.Get send POST request
func Post(c context.Context, uri string, params url.Values) (body []byte, err error) {
	req, err := NewRequest("POST", uri, params)
	if err != nil {
		return
	}
	body, err = Do(c, req)
	return
}

func NewUploadRequest(link string, params map[string]string, name, path string) (*http.Request, error) {
	//打开文件句柄操作
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error opening file")
		return nil, err
	}
	defer file.Close()

	//创建一个模拟的form中的一个选项,这个form项现在是空的
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作, 设置文件的上传参数叫uploadfile, 文件名是filename,
	//相当于现在还没选择文件, form项里选择文件的选项
	fileWriter, err := bodyWriter.CreateFormFile("file", path)
	if err != nil {
		fmt.Println("error writing to buffer")
		return nil, err
	}

	//iocopy 这里相当于选择了文件,将文件放到form中
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}

	//获取上传文件的类型,multipart/form-data; boundary=...
	contentType := bodyWriter.FormDataContentType()

	//这里就是上传的其他参数设置,可以使用 bodyWriter.WriteField(key, val) 方法
	//也可以自己在重新使用  multipart.NewWriter 重新建立一项,这个再server 会有例子
	//这种设置值得仿佛 和下面再从新创建一个的一样
	for key, val := range params {
		e := bodyWriter.WriteField(key, val)
		if e != nil {
			return nil, e
		}
	}

	//这个很关键,必须这样写关闭,不能使用defer关闭,不然会导致错误
	bodyWriter.Close()

	//发送post请求到服务端
	request, err := http.NewRequest(http.MethodPost, link, bodyBuf)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", contentType)
	return request, err

}
