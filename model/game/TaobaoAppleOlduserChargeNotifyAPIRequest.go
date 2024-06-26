package game

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAppleOlduserChargeNotifyAPIRequest 老用户激活并兑换通知接口 API请求
// taobao.apple.olduser.charge.notify
//
// 老用户激活并兑换通知接口
type TaobaoAppleOlduserChargeNotifyAPIRequest struct {
	model.Params
	// 结果code
	_resultCode string
	// 结果信息说明
	_resultMsg string
	// 业务参数
	_mainData *AppleTopOldSignNotifyDo
}

// NewTaobaoAppleOlduserChargeNotifyRequest 初始化TaobaoAppleOlduserChargeNotifyAPIRequest对象
func NewTaobaoAppleOlduserChargeNotifyRequest() *TaobaoAppleOlduserChargeNotifyAPIRequest {
	return &TaobaoAppleOlduserChargeNotifyAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAppleOlduserChargeNotifyAPIRequest) Reset() {
	r._resultCode = ""
	r._resultMsg = ""
	r._mainData = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAppleOlduserChargeNotifyAPIRequest) GetApiMethodName() string {
	return "taobao.apple.olduser.charge.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAppleOlduserChargeNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAppleOlduserChargeNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResultCode is ResultCode Setter
// 结果code
func (r *TaobaoAppleOlduserChargeNotifyAPIRequest) SetResultCode(_resultCode string) error {
	r._resultCode = _resultCode
	r.Set("result_code", _resultCode)
	return nil
}

// GetResultCode ResultCode Getter
func (r TaobaoAppleOlduserChargeNotifyAPIRequest) GetResultCode() string {
	return r._resultCode
}

// SetResultMsg is ResultMsg Setter
// 结果信息说明
func (r *TaobaoAppleOlduserChargeNotifyAPIRequest) SetResultMsg(_resultMsg string) error {
	r._resultMsg = _resultMsg
	r.Set("result_msg", _resultMsg)
	return nil
}

// GetResultMsg ResultMsg Getter
func (r TaobaoAppleOlduserChargeNotifyAPIRequest) GetResultMsg() string {
	return r._resultMsg
}

// SetMainData is MainData Setter
// 业务参数
func (r *TaobaoAppleOlduserChargeNotifyAPIRequest) SetMainData(_mainData *AppleTopOldSignNotifyDo) error {
	r._mainData = _mainData
	r.Set("main_data", _mainData)
	return nil
}

// GetMainData MainData Getter
func (r TaobaoAppleOlduserChargeNotifyAPIRequest) GetMainData() *AppleTopOldSignNotifyDo {
	return r._mainData
}

var poolTaobaoAppleOlduserChargeNotifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAppleOlduserChargeNotifyRequest()
	},
}

// GetTaobaoAppleOlduserChargeNotifyRequest 从 sync.Pool 获取 TaobaoAppleOlduserChargeNotifyAPIRequest
func GetTaobaoAppleOlduserChargeNotifyAPIRequest() *TaobaoAppleOlduserChargeNotifyAPIRequest {
	return poolTaobaoAppleOlduserChargeNotifyAPIRequest.Get().(*TaobaoAppleOlduserChargeNotifyAPIRequest)
}

// ReleaseTaobaoAppleOlduserChargeNotifyAPIRequest 将 TaobaoAppleOlduserChargeNotifyAPIRequest 放入 sync.Pool
func ReleaseTaobaoAppleOlduserChargeNotifyAPIRequest(v *TaobaoAppleOlduserChargeNotifyAPIRequest) {
	v.Reset()
	poolTaobaoAppleOlduserChargeNotifyAPIRequest.Put(v)
}
