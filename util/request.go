package util

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ucanme/fastgo/library/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type ContentType string

const (
	ContentType_JSON ContentType = "application/x-www-form-urlencoded"
	ContentType_XML  ContentType = "application/json"
)

func Get(reqUrl string, reqParams map[string]string) ([]byte, error) {
	params := url.Values{}
	urlPath, _ := url.Parse(reqUrl)
	for key, val := range reqParams {
		params.Set(key, val)
	}

	urlPath.RawQuery = params.Encode()
	resp, err := http.Get(urlPath.String())
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
func Post(url string, body []byte, params map[string]string, headers map[string]string) ([]byte, error) {
	//add post body

	var req *http.Request
	fmt.Println(string(body))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.LogError(map[string]interface{}{"url":url,"err":err})
		return nil, errors.New("new request is fail: %v \n")
	}
	req.Header.Set("Content-type", "application/json")
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{Timeout: 10 * time.Second}
	resp,err :=  client.Do(req)
	if err != nil{
		log.LogError(map[string]interface{}{"url":url,"err":err})
		return nil,err
	}
	data,err := ioutil.ReadAll(resp.Body)
	log.LogNotice(map[string]interface{}{"url":req.URL.String(),"resp":string(data),"err":err})
	return data,err
}
