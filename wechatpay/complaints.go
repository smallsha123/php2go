package wechatPay

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/smallsha123/php2go/errorx"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"io/ioutil"
	"log"
)

// ComplaintNotificationsUrlQuery 商户投诉地址查询
func (s *WechatPay) ComplaintNotificationsUrlQuery() (resp *MerchantComplaintsUrlReply, err error) {
	if s.Client == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}
	result, err := s.Client.Get(s.Ctx, "https://api.mch.weixin.qq.com/v3/merchant-service/complaint-notifications")
	if err != nil {
		if core.IsAPIError(err, "INVALID_REQUEST") {
			s.Logger.Errorf("load merchant private key error %v", err)
		}
		// 处理的其他错误
		s.Logger.Errorf("load merchant private key error %v", err)
	}
	body, _ := ioutil.ReadAll(result.Response.Body)
	_ = json.Unmarshal(body, &resp)
	return
}

// ComplaintNotificationsUrlCreate 商户投诉地址创建
func (s *WechatPay) ComplaintNotificationsUrlCreate() (resp *MerchantComplaintsUrlReply, err error) {
	if s.Client == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}
	data := MerchantComplaintsUrlCreateReq{}
	data.URL = "https://www.baidu.com"
	post, _ := s.Client.Post(s.Ctx, "https://api.mch.weixin.qq.com/v3/merchant-service/complaint-notifications", data)
	if post == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "request failed.")
	}

	var respErr map[string]string

	body, _ := ioutil.ReadAll(post.Response.Body)
	_ = json.Unmarshal(body, &respErr)

	// fmt.Printf("%+v", respErr)
	if _, ok := respErr["code"]; ok {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, respErr["message"])
	}
	_ = json.Unmarshal(body, &resp)
	return
}

// ComplaintNotificationsUrlUpdate 商户投诉地址更新
func (s *WechatPay) ComplaintNotificationsUrlUpdate(url string) (resp *MerchantComplaintsUrlReply, err error) {
	if s.Client == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}
	data := MerchantComplaintsUrlCreateReq{}
	data.URL = url
	post, _ := s.Client.Put(s.Ctx, "https://api.mch.weixin.qq.com/v3/merchant-service/complaint-notifications", data)
	if post == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "request failed.")
	}

	var respErr map[string]string

	body, _ := ioutil.ReadAll(post.Response.Body)
	_ = json.Unmarshal(body, &respErr)

	// fmt.Printf("%+v", respErr)
	if _, ok := respErr["code"]; ok {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, respErr["message"])
	}
	_ = json.Unmarshal(body, &resp)
	return
}

// ComplaintNotificationsUrlDelete 商户投诉地址删除
func (s *WechatPay) ComplaintNotificationsUrlDelete() (resp interface{}, err error) {
	if s.Client == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}
	post, err := s.Client.Delete(s.Ctx, "https://api.mch.weixin.qq.com/v3/merchant-service/complaint-notifications", "")
	if err != nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, err.Error())
	}
	if post == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "request failed.")
	}

	var respErr map[string]string

	body, _ := ioutil.ReadAll(post.Response.Body)
	_ = json.Unmarshal(body, &respErr)

	fmt.Printf("%+v", respErr)
	if _, ok := respErr["code"]; ok {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, respErr["message"])
	}
	return
}

// ComplaintInfo 投诉详情
func (s *WechatPay) ComplaintInfo(complaintId string) (resp *MerchantComplaintsInfo, err error) {
	if s.Client == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}
	post, _ := s.Client.Get(s.Ctx, "https://api.mch.weixin.qq.com/v3/merchant-service/complaints-v2/"+complaintId)
	if post == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "request failed.")
	}
	body, _ := ioutil.ReadAll(post.Response.Body)

	var respErr map[string]string
	_ = json.Unmarshal(body, &respErr)
	if _, ok := respErr["code"]; ok {
		fmt.Printf("返回信息：%+v；投诉单号：%s", respErr, complaintId)
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, respErr["message"])
	}
	_ = json.Unmarshal(body, &resp)
	phone, err := utils.DecryptOAEP(resp.PayerPhone, s.MchPrivateKey)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "手机号解密失败："+err.Error())
	}
	resp.PayerPhone = phone
	return
}

// DoAEADOpen 回调报文解密
func (s *WechatPay) DoAEADOpen(nonce, ciphertext, additionalData string) (plaintText *MerchantComplaintsNotifyPlaintext, err error) {
	c, err := aes.NewCipher([]byte(s.WechatPayParams.MchAPIv3Key))
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	info, err := aesgcm.Open(nil, []byte(nonce), data, []byte(additionalData))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(info, &plaintText)
	if err != nil {
		return nil, err
	}
	return
}

func (s WechatPay) GetComplaintMedia() (resp *MerchantComplaintsInfo, err error) {
	if s.Client == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "实例化错误")
	}
	post, _ := s.Client.Get(s.Ctx, "https://api.mch.weixin.qq.com/v3/merchant-service/images/ChsyMDAwMDAwMjAyMzAzMjAxMzAwNzAxMTY0MDEYACDM5d%2BgBigBMAE4AQ%3D%3D")
	if post == nil {
		return nil, errorx.NewCodeError(errorx.ERR_DEFAULT, "request failed.")
	}
	// s.Logger.Infof("信息：%v", post.Response.Body)
	all, err := ioutil.ReadAll(post.Response.Body)
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile("image.jpg", all, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return
}
