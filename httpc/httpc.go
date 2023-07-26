package httpc

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/smallsha123/php2go/errorx"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpc"
	"io"
	"io/ioutil"
	"net/http"
)

func BaseResponse(ctx context.Context, url string, data interface{}, header ...http.Header) (map[string]interface{}, error) {
	logger := logx.WithContext(ctx)

	var reader io.Reader
	if data != nil {
		b, _ := json.Marshal(data)
		reader = bytes.NewReader(b)
	}

	r, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	for _, h := range header {
		for k, listStr := range h {
			for _, s := range listStr {
				r.Header.Set(k, s)
			}
		}
	}

	resp, err := httpc.DoRequest(r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var j map[string]interface{}
	err = json.Unmarshal(b, &j)

	logger.Infof("url:%s respData:%+v", url, j)

	return j, err
}

func Post(ctx context.Context, url string, data interface{}, header ...http.Header) (interface{}, error) {
	logger := logx.WithContext(ctx)

	j, err := BaseResponse(ctx, url, data, header...)
	if err != nil {
		logger.Errorf("request failed. err:%s url:%s data:%+v j:%v", err.Error(), url, data, j)
		return nil, err
	}

	if _, ok := j["code"]; !ok {
		logger.Errorf("not found the key of code. url:%s data:%+v j:%v", url, data, j)
		return nil, errors.New("not found the key of code")
	}

	if j["code"].(float64) != errorx.OK {
		logger.Errorf("code validate failed. url:%s data:%+v j:%v", url, data, j)
		return nil, errorx.NewCodeError(cast.ToInt(j["code"]), cast.ToString(j["message"]))
	}

	return j["data"], nil
}

func Get(ctx context.Context, url string, header ...http.Header) (interface{}, error) {
	logger := logx.WithContext(ctx)
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for _, h := range header {
		for k, listStr := range h {
			for _, s := range listStr {
				r.Header.Set(k, s)
			}
		}
	}

	resp, err := httpc.DoRequest(r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var j map[string]interface{}
	err = json.Unmarshal(b, &j)
	if err != nil {
		return nil, err
	}

	if _, ok := j["code"]; !ok {
		logger.Errorf("not found the key of code. url:%s j:%v", url, j)
		return nil, errors.New("not found the key of code")
	}

	if j["code"].(float64) != errorx.OK {
		logger.Errorf("code validate failed. url:%s j:%v", url, j)
		return nil, errorx.NewCodeError(cast.ToInt(j["code"]), cast.ToString(j["message"]))
	}

	return j["data"], nil
}
