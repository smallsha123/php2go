package ctxdata

import (
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "jwtUserId"
var CtxKeyJwtForm = "jwtForm"
var CtxKeyJwtUserInfo = "jwtUserInfo"
var CtxKeyPermission = "dataPermission"

type UserinfoReply struct {
	Id           int64  `json:"id"`
	OpenId       string `json:"open_id"`
	UnionId      string `json:"union_id"`
	AppId        string `json:"app_id"`
	Account      string `json:"account"`
	Realname     string `json:"realname"`
	Gender       string `json:"gender"`
	Avatar       string `json:"avatar"`
	Phone        string `json:"phone"`
	Role         int64  `json:"role"`
	PlatformType int64  `json:"platform_type"`
	CompanyId    int64  `json:"company_id"`
	ShopId       int64  `json:"shop_id"`
}

type UserMapReply struct {
	Key      int64  `json:"key"`
	Realname string `json:"realname"`
}

type MenuPermissionReply struct {
	DataPermission int64           `json:"data_permission"`
	UserStr        string          `json:"user_str"`
	UserList       []*UserMapReply `json:"user_list"`
}

func GetUidFromCtx(ctx context.Context) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	}
	return uid
}

func GetUserInfoFromCtx(ctx context.Context) UserinfoReply {
	userMap := ctx.Value(CtxKeyJwtUserInfo)
	b, _ := json.Marshal(userMap)
	var resp UserinfoReply
	_ = json.Unmarshal(b, &resp)

	return resp
}

// GetDataPermissionFromCtx 数据权限
func GetDataPermissionFromCtx(ctx context.Context) MenuPermissionReply {
	m := ctx.Value(CtxKeyPermission)
	b, _ := json.Marshal(m)
	var resp MenuPermissionReply
	_ = json.Unmarshal(b, &resp)

	return resp
}
