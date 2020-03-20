/**
 * @author: wuji
 * @file:  http
 * @desc:
 * @date: 2019/11/15 4:46 下午
 */

package httputil

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
func Post(url string, body []byte, params map[string]string, headers map[string]string) (*http.Response, error) {
	//add post body

	var req *http.Request
	fmt.Println(string(body))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
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
	client := &http.Client{}
	log.Printf("Go POST URL : %s \n", req.URL.String())
	return client.Do(req)
}
