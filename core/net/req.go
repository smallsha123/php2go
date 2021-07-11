package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Result string "json"
	Error  error
}

func NewResponse(result string, error error) *Response {
	return &Response{Result: result, Error: error}
}

func (s *Response) SetResult(result string, error error) {
	s.Result = result
	s.Error = error
}

func (s Response) GetResult() {

}

//Get http get method
func Get(url string, params map[string]string, headers map[string]string) ([]byte, error) {
	//new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("new request is fail ")
	}
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
	log.Printf("Go GET URL : %s \n", req.URL.String())

	//发送请求
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close() //一定要关闭res.Body
	//读取body
	resBody, err := ioutil.ReadAll(res.Body) //把  body 内容读入字符串 s
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

//Post http post method
func Post(url string, body map[string]interface{}, params map[string]string, headers map[string]string) ([]byte, error) {
	//add post body
	var bodyJson []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			log.Println(err)
			return nil, errors.New("http post body to json failed")
		}
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJson))
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

	//发送请求
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close() //一定要关闭res.Body
	//读取body
	resBody, err := ioutil.ReadAll(res.Body) //把  body 内容读入字符串 s
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
