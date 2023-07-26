package applet

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"io/ioutil"
	"time"

	"github.com/tidwall/gjson"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/kr/pretty"
	"github.com/smallsha123/php2go/errorx"
	"github.com/smallsha123/php2go/httpc"
)

var PowerWechatMap map[string]PowerWechat

type PowerWechat struct {
	MiniProgramApp    *miniProgram.MiniProgram
	PowerWechatParams *PowerWechatParams
	AccessToken       string
	Ctx               context.Context
	logx.Logger
}

// InitPowerWechatMap 初始化小程序信息
func InitPowerWechatMap(powerWechatParams []PowerWechatParams) {

	PowerWechatMap = make(map[string]PowerWechat)

	for _, powerWechatParam := range powerWechatParams {

		miniProgramApp, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
			AppID:  powerWechatParam.AppID,  // 小程序appid
			Secret: powerWechatParam.Secret, // 小程序app secret
			Token:  powerWechatParam.Token,  // 小程序Token
			AESKey: powerWechatParam.AESKey, // 小程序AESKey

			HttpDebug: powerWechatParam.HttpDebug,
			// Log: miniProgram.Log{
			//	Level: params.Level,
			//	File:  params.File,
			// },
			// 可选，不传默认走程序内存
			Cache: NewRedisCache(powerWechatParam.Redis),
		})
		if err != nil {
			logx.Errorf("【实例化小程序错误】err:%+v", err.Error())
		}

		PowerWechatMap[powerWechatParam.AppID] = PowerWechat{
			MiniProgramApp:    miniProgramApp,
			PowerWechatParams: &powerWechatParam,
		}
	}
}

// NewPowerWechat 实例化小程序
func NewPowerWechat(params *PowerWechatParams, ctx context.Context) (power *PowerWechat, err error) {

	// 基于上下文的日志链路
	logger := logx.WithContext(ctx)

	// 判断小程序是否存在
	client, ok := PowerWechatMap[params.AppID]
	if ok == false {
		logger.Errorf("【小程序实例化】小程序信息不存在：%+v", params)
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "小程序信息不存在Appid: "+params.AppID)
	}

	// 获取小程序accessToken
	token, err := client.MiniProgramApp.AccessToken.GetToken(false)
	accessToken := ""
	if err != nil {
		logger.Errorf("【小程序实例化】获取accessToken错误 failed. err:%+v", err.Error())
		//return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, err.Error())
	}
	if token != nil {
		accessToken = token.AccessToken
	}
	return &PowerWechat{
		MiniProgramApp:    client.MiniProgramApp,
		PowerWechatParams: client.PowerWechatParams,
		AccessToken:       accessToken,
		Ctx:               ctx,
		Logger:            logger,
	}, nil
}

// GetAccessToken 获取小程序accessToken
func (w *PowerWechat) GetAccessToken() string {
	token, err := w.MiniProgramApp.AccessToken.GetToken(false)
	if err != nil {
		w.Errorf("【获取小程序accessToken】获取accessToken错误 failed. err:%+v", err.Error())
		return ""
	}
	return token.AccessToken
}

// RegisterCustomerService 创建小程序客服子商户
func (w *PowerWechat) RegisterCustomerService(req *RegisterCustomerServiceReq) (resp *RegisterCustomerServiceReply, err error) {
	a := &httpc.APIResult{}
	url := "https://api.weixin.qq.com/cgi-bin/business/register?access_token=" + w.AccessToken
	post, err := a.Post(w.Ctx, url, req)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(post.Response.Body)
	fmt.Printf("%s", string(body))
	response := string(body)
	businessId := gjson.Get(response, "business_id").Int()
	if businessId > 0 {
		serviceReply := RegisterCustomerServiceReply{BusinessId: businessId}
		return &serviceReply, nil
	} else {
		errMsg := gjson.Get(response, "errmsg").String()
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, errMsg)
	}
}

// 回复投诉内容
func (w *PowerWechat) ReplyComplaintService(token string, req *RegisterMiniServiceReq) (resp *WxErrCodeReply, err error) {
	a := &httpc.APIResult{}
	url := "https://api.weixin.qq.com/wxaapi/minishop/bussiRespondComplaint?access_token=" + token

	postData := map[string]interface{}{}

	postData["complaintOrderId"] = req.ComplaintOrderId
	if req.Content != "" {
		postData["content"] = req.Content
	}
	if req.MediaIdList != nil {
		postData["mediaIdList"] = req.MediaIdList
	}
	if req.BussiHandle != 0 {
		postData["bussiHandle"] = req.BussiHandle
	}

	fmt.Printf("%# v\n", pretty.Formatter(postData))
	post, err := a.Post(w.Ctx, url, postData)

	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(post.Response.Body)

	errcode := &WxErrCodeReply{}
	errs := json.Unmarshal(body, errcode)

	if err != nil {
		return nil, errs
	}

	return errcode, nil
}

// 获取投诉详情
func (w *PowerWechat) ReplyInfoService(token string, order_id string) (resp *WxInfoReply, err error) {
	a := &httpc.APIResult{}
	url := "https://api.weixin.qq.com/wxaapi/minishop/complaintOrderDetail?complaintOrderId=%s&access_token=%s"
	url = fmt.Sprintf(url, order_id, token)

	post, err := a.Get(w.Ctx, url)

	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(post.Response.Body)

	resp = &WxInfoReply{}
	errs := json.Unmarshal(body, resp)
	if err != nil {
		return nil, errs
	}

	return resp, nil
}

type RedisCache struct {
	redisClient *redis.Redis
}

func NewRedisCache(client *redis.Redis) *RedisCache {
	return &RedisCache{redisClient: client}
}

func (r *RedisCache) Get(key string, defaultValue interface{}) (ptrValue interface{}, err error) {
	b, err := r.redisClient.Get(key)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(b), &ptrValue)
	return ptrValue, err
}

func (r *RedisCache) Set(key string, value interface{}, expires time.Duration) error {
	mValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.redisClient.Setex(key, string(mValue), cast.ToInt(expires.Seconds()))
	return err
}

func (r *RedisCache) Has(key string) bool {
	value, err := r.redisClient.Get(key)
	if value != "" && err == nil {
		return true
	}

	return false
}

func (r *RedisCache) AddNX(key string, value interface{}, ttl time.Duration) bool {
	mValue, err := json.Marshal(value)
	if err != nil {
		return false
	}

	b, err := r.redisClient.SetnxEx(key, string(mValue), cast.ToInt(ttl.Seconds()))
	if err != nil {
		fmt.Errorf("SetNX error: %+v \r\n", err)
	}
	return b
}

func (r *RedisCache) Add(key string, value interface{}, ttl time.Duration) (err error) {
	var obj interface{}
	obj, err = r.Get(key, obj)

	mValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.redisClient.Setex(key, string(mValue), cast.ToInt(ttl.Seconds()))
}

func (r *RedisCache) Remember(key string, ttl time.Duration, callback func() (interface{}, error)) (obj interface{}, err error) {
	var value interface{}
	value, err = r.Get(key, value)

	// If the item exists in the cache we will just return this immediately and if
	// not we will execute the given Closure and cache the result of that for a
	// given number of seconds so it's available for all subsequent requests.
	if err != nil {
		return nil, err
	} else if value != nil {
		return value, err
	}

	value, err = callback()
	if err != nil {
		return nil, err
	}

	result := r.Put(key, value, ttl)
	if !result {
		err = errors.New(fmt.Sprintf("remember cache put err, ttl:%d", ttl))
	}
	// ErrCacheMiss and query value from source
	return value, err
}

func (r *RedisCache) Put(key interface{}, value interface{}, ttl time.Duration) bool {
	err := r.Set(key.(string), value, ttl)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true

}
