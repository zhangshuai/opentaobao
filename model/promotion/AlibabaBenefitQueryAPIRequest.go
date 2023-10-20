package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibababenefitqueryAPIRequest 奖池奖品查询列表 API请求
// alibaba.benefit.query
//
// 功能：奖池奖品查询列表
// 业务逻辑：程序中通过奖池编号ename,业务方身份appName来查询奖池提供的奖品返回给
// 小程。
// 安全保障：为保证数据不会越权，需要卖家授，并且验证系统参数appKey。只有通过授权的，并且
// appkey验证通过的，才会查数据 并透出，否则直接失败。
// 因为appkey是系统参数，并且程序内部可以验证appkey和业务身份appName的关系
// 是否一致，所以可以保证参数appName的合法性，没有越权。
type AlibababenefitqueryAPIRequest struct {
	model.Params
	// 奖池编号
	_ename string
	// 表示奖池类型（发奖奖池传1，抽奖传0）
	_awardType string
	// 商家来源身份标识（"promotion-"+appId）
	_appName string
}

// NewAlibababenefitqueryRequest 初始化AlibababenefitqueryAPIRequest对象
func NewAlibababenefitqueryRequest() *AlibababenefitqueryAPIRequest {
	return &AlibababenefitqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibababenefitqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.benefit.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibababenefitqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibababenefitqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEname is Ename Setter
// 奖池编号
func (r *AlibababenefitqueryAPIRequest) SetEname(_ename string) error {
	r._ename = _ename
	r.Set("ename", _ename)
	return nil
}

// GetEname Ename Getter
func (r AlibababenefitqueryAPIRequest) GetEname() string {
	return r._ename
}

// SetAwardType is AwardType Setter
// 表示奖池类型（发奖奖池传1，抽奖传0）
func (r *AlibababenefitqueryAPIRequest) SetAwardType(_awardType string) error {
	r._awardType = _awardType
	r.Set("award_type", _awardType)
	return nil
}

// GetAwardType AwardType Getter
func (r AlibababenefitqueryAPIRequest) GetAwardType() string {
	return r._awardType
}

// SetAppName is AppName Setter
// 商家来源身份标识（&#34;promotion-&#34;+appId）
func (r *AlibababenefitqueryAPIRequest) SetAppName(_appName string) error {
	r._appName = _appName
	r.Set("app_name", _appName)
	return nil
}

// GetAppName AppName Getter
func (r AlibababenefitqueryAPIRequest) GetAppName() string {
	return r._appName
}
