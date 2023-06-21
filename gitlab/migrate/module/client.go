package module

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"strconv"
)

func GetHttps(url, token string) (*http.Response, error) {

	// 创建各类对象
	var client *http.Client
	var request *http.Request
	var err error
	client = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}

	// 获取 request请求
	request, err = http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("GetHttpSkip Request Error:", err)
		return nil, nil
	}

	// 加入 token
	request.Header.Add("PRIVATE-TOKEN", token)
	resp, err := client.Do(request)
	if err != nil {
		log.Println("GetHttpSkip Response Error:", err)
		return nil, nil
	}
	defer client.CloseIdleConnections()
	return resp, nil
}

func PostHttp(url, token string, reader io.Reader) (*http.Response, error) {

	// 创建各类对象
	var client *http.Client
	var request *http.Request
	var err error

	client = &http.Client{}

	// 获取 request请求
	request, err = http.NewRequest("POST", url, reader)

	if err != nil {
		log.Println("PostHttpSkip Request Error:", err)
		return nil, nil
	}
	request.Header.Set("Content-Type", "application/json")
	// 加入 token
	request.Header.Add("PRIVATE-TOKEN", token)

	resp, err := client.Do(request)
	if err != nil {
		log.Println("PostHttpSkip Response Error:", err)
		return nil, nil
	}
	defer client.CloseIdleConnections()
	return resp, nil
}

func GetBody(url, token string) ([]byte, error) {
	var body []byte

	r, err := GetHttps(url, token)
	if err != nil {
		log.Println("Access"+url+"Response Error:", err)
		return nil, nil
	}
	body, _ = io.ReadAll(r.Body)
	return body, nil

}
func GetPagestotal(url, token string) (int, error) {

	r, err := GetHttps(url, token)

	if err != nil {
		log.Println("Access"+url+"Response Error:", err)
		return 0, nil
	}
	totalPages, _ := strconv.Atoi(r.Header.Get("X-Total-Pages"))
	return totalPages, nil
}
