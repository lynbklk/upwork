package request

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func MakeCookies(cookies string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", cookies)
	request := http.Request{Header: header}
	return request.Cookies()
}

func MakeRequest(method, url, body string, params map[string]string, headers map[string]string, cookies []*http.Cookie) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}

	// header
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// cookies
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	// params
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	return req, nil
}

func DoWithTimeout(req *http.Request, timeout int, len int64) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, "Metrics", req.URL.Host)
	req = req.WithContext(ctx)
	req.Close = true

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(&io.LimitedReader{R: resp.Body, N: len})
}
