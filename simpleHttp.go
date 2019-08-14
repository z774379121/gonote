package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
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
