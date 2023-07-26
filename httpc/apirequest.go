package httpc

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"time"
)

// 常用 ContentType
const (
	ApplicationJSON = "application/json"
	ApplicationForm = "application/x-www-form-urlencoded; charset=UTF-8"
	ImageJPG        = "image/jpg"
	ImagePNG        = "image/png"
	VideoMP4        = "video/mp4"
)

var (
	regJSONTypeCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	regXMLTypeCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// APIResult Api请求返回
type APIResult struct {
	httpClient *http.Client
	// 本次请求所使用的 HTTPRequest
	Request *http.Request
	// 本次请求所获得的 HTTPResponse
	Response *http.Response
}

// Get 发送一个 HTTP Get 请求
func (a *APIResult) Get(ctx context.Context, requestURL string) (*APIResult, error) {
	return a.doRequest(ctx, http.MethodGet, requestURL, nil, ApplicationJSON, nil)
}

// Post 发送一个 HTTP Post 请求
func (a *APIResult) Post(ctx context.Context, requestURL string, requestBody interface{}) (*APIResult, error) {
	return a.requestWithJSONBody(ctx, http.MethodPost, requestURL, requestBody)
}

func (a *APIResult) PostForm(ctx context.Context, requestURL string, requestBody interface{}, header http.Header) (*APIResult, error) {
	reqBody, err := setBody(requestBody, ApplicationJSON)
	if err != nil {
		return nil, err
	}

	return a.doRequest(ctx, http.MethodPost, requestURL, header, ApplicationForm, reqBody)
}

// 请求json体
func (a *APIResult) requestWithJSONBody(ctx context.Context, method, requestURL string, body interface{}) (
	*APIResult, error,
) {
	reqBody, err := setBody(body, ApplicationJSON)
	if err != nil {
		return nil, err
	}

	return a.doRequest(ctx, method, requestURL, nil, ApplicationJSON, reqBody)
}

// setBody Set Request body from an interface
//
//revive:disable-next-line:cyclomatic 本函数实现需要考虑多种情况，但理解起来并不复杂，进行圈复杂度豁免
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	bodyBuf = &bytes.Buffer{}

	switch b := body.(type) {
	case string:
		_, err = bodyBuf.WriteString(b)
	case *string:
		_, err = bodyBuf.WriteString(*b)
	case []byte:
		_, err = bodyBuf.Write(b)
	case **os.File:
		_, err = bodyBuf.ReadFrom(*b)
	case io.Reader:
		_, err = bodyBuf.ReadFrom(b)
	default:
		if regJSONTypeCheck.MatchString(contentType) {
			err = json.NewEncoder(bodyBuf).Encode(body)
		} else if regXMLTypeCheck.MatchString(contentType) {
			err = xml.NewEncoder(bodyBuf).Encode(body)
		}
		// format.
	}
	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

func (a *APIResult) doRequest(
	ctx context.Context,
	method string,
	requestURL string,
	header http.Header,
	contentType string,
	reqBody io.Reader,
) (*APIResult, error) {

	var (
		err     error
		request *http.Request
	)

	// Construct Request
	if request, err = http.NewRequestWithContext(ctx, method, requestURL, reqBody); err != nil {
		return nil, err
	}

	// Header Setting Priority:
	// Fixed Headers > Per-Request Header Parameters

	// Add Request Header Parameters
	for key, values := range header {
		for _, v := range values {
			request.Header.Add(key, v)
		}
	}

	// Set Fixed Headers
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Content-Type", contentType)
	ua := fmt.Sprintf("TkPlatform/%s (%s) GO/%s", "1.0.0", runtime.GOOS, runtime.Version())
	request.Header.Set("User-Agent", ua)

	// Send HTTP Request
	result, err := a.doHTTP(request)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (a *APIResult) doHTTP(req *http.Request) (result *APIResult, err error) {
	result = &APIResult{
		Request: req,
	}
	if a.httpClient == nil {
		a.httpClient = &http.Client{
			Timeout: 330 * time.Second,
		}
	}
	result.Response, err = a.httpClient.Do(req)
	return result, err
}
