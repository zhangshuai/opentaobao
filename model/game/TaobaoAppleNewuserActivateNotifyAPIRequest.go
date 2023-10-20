package game

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoapplenewuseractivatenotifyAPIRequest 新用户激活通知接口 API请求
// taobao.apple.newuser.activate.notify
//
// 资和信主动通知激活结果
type TaobaoapplenewuseractivatenotifyAPIRequest struct {
	model.Params
	// 结果对应值，00位成功，其他为失败
	_resultCode string
	// 处理结果中文描述
	_resultMsg string
	// 主业务参数
	_mainData *AppleTopActivateNotifyDo
}

// NewTaobaoapplenewuseractivatenotifyRequest 初始化TaobaoapplenewuseractivatenotifyAPIRequest对象
func NewTaobaoapplenewuseractivatenotifyRequest() *TaobaoapplenewuseractivatenotifyAPIRequest {
	return &TaobaoapplenewuseractivatenotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoapplenewuseractivatenotifyAPIRequest) GetApiMethodName() string {
	return "taobao.apple.newuser.activate.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoapplenewuseractivatenotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoapplenewuseractivatenotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResultCode is ResultCode Setter
// 结果对应值，00位成功，其他为失败
func (r *TaobaoapplenewuseractivatenotifyAPIRequest) SetResultCode(_resultCode string) error {
	r._resultCode = _resultCode
	r.Set("result_code", _resultCode)
	return nil
}

// GetResultCode ResultCode Getter
func (r TaobaoapplenewuseractivatenotifyAPIRequest) GetResultCode() string {
	return r._resultCode
}

// SetResultMsg is ResultMsg Setter
// 处理结果中文描述
func (r *TaobaoapplenewuseractivatenotifyAPIRequest) SetResultMsg(_resultMsg string) error {
	r._resultMsg = _resultMsg
	r.Set("result_msg", _resultMsg)
	return nil
}

// GetResultMsg ResultMsg Getter
func (r TaobaoapplenewuseractivatenotifyAPIRequest) GetResultMsg() string {
	return r._resultMsg
}

// SetMainData is MainData Setter
// 主业务参数
func (r *TaobaoapplenewuseractivatenotifyAPIRequest) SetMainData(_mainData *AppleTopActivateNotifyDo) error {
	r._mainData = _mainData
	r.Set("main_data", _mainData)
	return nil
}

// GetMainData MainData Getter
func (r TaobaoapplenewuseractivatenotifyAPIRequest) GetMainData() *AppleTopActivateNotifyDo {
	return r._mainData
}
