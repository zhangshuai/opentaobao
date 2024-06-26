package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopSdkFeedbackUploadAPIRequest sdk信息回调 API请求
// taobao.top.sdk.feedback.upload
//
// sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
type TaobaoTopSdkFeedbackUploadAPIRequest struct {
	model.Params
	// 1、回传加密信息
	_type string
	// 具体内容，json形式
	_content string
}

// NewTaobaoTopSdkFeedbackUploadRequest 初始化TaobaoTopSdkFeedbackUploadAPIRequest对象
func NewTaobaoTopSdkFeedbackUploadRequest() *TaobaoTopSdkFeedbackUploadAPIRequest {
	return &TaobaoTopSdkFeedbackUploadAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTopSdkFeedbackUploadAPIRequest) Reset() {
	r._type = ""
	r._content = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopSdkFeedbackUploadAPIRequest) GetApiMethodName() string {
	return "taobao.top.sdk.feedback.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopSdkFeedbackUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopSdkFeedbackUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// 1、回传加密信息
func (r *TaobaoTopSdkFeedbackUploadAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoTopSdkFeedbackUploadAPIRequest) GetType() string {
	return r._type
}

// SetContent is Content Setter
// 具体内容，json形式
func (r *TaobaoTopSdkFeedbackUploadAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaoTopSdkFeedbackUploadAPIRequest) GetContent() string {
	return r._content
}

var poolTaobaoTopSdkFeedbackUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTopSdkFeedbackUploadRequest()
	},
}

// GetTaobaoTopSdkFeedbackUploadRequest 从 sync.Pool 获取 TaobaoTopSdkFeedbackUploadAPIRequest
func GetTaobaoTopSdkFeedbackUploadAPIRequest() *TaobaoTopSdkFeedbackUploadAPIRequest {
	return poolTaobaoTopSdkFeedbackUploadAPIRequest.Get().(*TaobaoTopSdkFeedbackUploadAPIRequest)
}

// ReleaseTaobaoTopSdkFeedbackUploadAPIRequest 将 TaobaoTopSdkFeedbackUploadAPIRequest 放入 sync.Pool
func ReleaseTaobaoTopSdkFeedbackUploadAPIRequest(v *TaobaoTopSdkFeedbackUploadAPIRequest) {
	v.Reset()
	poolTaobaoTopSdkFeedbackUploadAPIRequest.Put(v)
}
