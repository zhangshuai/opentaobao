package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjMemberHasbindAPIRequest 喵街会员是否绑定 API请求
// alibaba.mj.member.hasbind
//
// 喵街检测用户是否为数字化会员
type AlibabaMjMemberHasbindAPIRequest struct {
	model.Params
	// open_id
	_openId string
	// user_id
	_userId int64
}

// NewAlibabaMjMemberHasbindRequest 初始化AlibabaMjMemberHasbindAPIRequest对象
func NewAlibabaMjMemberHasbindRequest() *AlibabaMjMemberHasbindAPIRequest {
	return &AlibabaMjMemberHasbindAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMjMemberHasbindAPIRequest) Reset() {
	r._openId = ""
	r._userId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjMemberHasbindAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.member.hasbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjMemberHasbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjMemberHasbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenId is OpenId Setter
// open_id
func (r *AlibabaMjMemberHasbindAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r AlibabaMjMemberHasbindAPIRequest) GetOpenId() string {
	return r._openId
}

// SetUserId is UserId Setter
// user_id
func (r *AlibabaMjMemberHasbindAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaMjMemberHasbindAPIRequest) GetUserId() int64 {
	return r._userId
}

var poolAlibabaMjMemberHasbindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMjMemberHasbindRequest()
	},
}

// GetAlibabaMjMemberHasbindRequest 从 sync.Pool 获取 AlibabaMjMemberHasbindAPIRequest
func GetAlibabaMjMemberHasbindAPIRequest() *AlibabaMjMemberHasbindAPIRequest {
	return poolAlibabaMjMemberHasbindAPIRequest.Get().(*AlibabaMjMemberHasbindAPIRequest)
}

// ReleaseAlibabaMjMemberHasbindAPIRequest 将 AlibabaMjMemberHasbindAPIRequest 放入 sync.Pool
func ReleaseAlibabaMjMemberHasbindAPIRequest(v *AlibabaMjMemberHasbindAPIRequest) {
	v.Reset()
	poolAlibabaMjMemberHasbindAPIRequest.Put(v)
}
