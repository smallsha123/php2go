package response

import (
	"net/http"
	"strconv"
	"time"

	"google.golang.org/grpc/status"

	"github.com/pkg/errors"
	"github.com/smallsha123/php2go/errorx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	TimeStamp int64       `json:"timestamp"`
}

// 统一封装成功响应值
func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		err = errors.Cause(err)
		switch e := err.(type) {
		case *errorx.CodeError: // 业务输出错误
			body.Code = e.Code
			body.Message = e.Message
			body.Data = e.Data
			body.TimeStamp = time.Now().Unix()
		default:
			if gstatus, ok := status.FromError(err); ok {
				// grpc err错误
				body.Code = errorx.ERR_DEFAULT
				body.Message = gstatus.Message()
				body.TimeStamp = time.Now().Unix()
			} else {
				// 系统未知错误
				body.Code = errorx.ERR_DEFAULT
				body.Message = errorx.MapErrMsg(body.Code)
				body.TimeStamp = time.Now().Unix()
			}
		}
	} else {
		body.Code = errorx.OK
		body.Message = errorx.MapErrMsg(body.Code)
		body.Data = resp
		body.TimeStamp = time.Now().Unix()
	}
	httpx.OkJson(w, body)
}

func Code(err error) int {
	if se, ok := status.FromError(err); ok {
		code, err := strconv.Atoi((se).Code().String())
		if err != nil {
			return -1
		}
		return code
	}
	return -1
}
