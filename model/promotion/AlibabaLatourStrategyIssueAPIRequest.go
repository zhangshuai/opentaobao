package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里巴巴权益发放接口 API请求
alibaba.latour.strategy.issue

阿里巴巴权益平台权益发放接口
*/
type AlibabaLatourStrategyIssueAPIRequest struct {
    model.Params
    // 扩展参数
    _extraData   string
    // 算法容灾
    _failoverAlgorithmResult   bool
    // 幂等id
    _idempotentId   string
    // 发放渠道
    _channel   string
    // 除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
    _userId   string
    // 转换用户类型
    _transformedUserType   string
    // 是否需要过安全
    _needIdentifyRisk   bool
    // 除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
    _userNick   string
    // 投放计划code
    _strategyCode   string
    // 用户类型
    _userType   string
    // 指定发放权益code
    _selectedBenefitCode   string
}

// 初始化AlibabaLatourStrategyIssueAPIRequest对象
func NewAlibabaLatourStrategyIssueRequest() *AlibabaLatourStrategyIssueAPIRequest{
    return &AlibabaLatourStrategyIssueAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLatourStrategyIssueAPIRequest) GetApiMethodName() string {
    return "alibaba.latour.strategy.issue"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLatourStrategyIssueAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtraData Setter
// 扩展参数
func (r *AlibabaLatourStrategyIssueAPIRequest) SetExtraData(_extraData string) error {
    r._extraData = _extraData
    r.Set("extra_data", _extraData)
    return nil
}

// ExtraData Getter
func (r AlibabaLatourStrategyIssueAPIRequest) GetExtraData() string {
    return r._extraData
}
// FailoverAlgorithmResult Setter
// 算法容灾
func (r *AlibabaLatourStrategyIssueAPIRequest) SetFailoverAlgorithmResult(_failoverAlgorithmResult bool) error {
    r._failoverAlgorithmResult = _failoverAlgorithmResult
    r.Set("failover_algorithm_result", _failoverAlgorithmResult)
    return nil
}

// FailoverAlgorithmResult Getter
func (r AlibabaLatourStrategyIssueAPIRequest) GetFailoverAlgorithmResult() bool {
    return r._failoverAlgorithmResult
}
// IdempotentId Setter
// 幂等id
func (r *AlibabaLatourStrategyIssueAPIRequest) SetIdempotentId(_idempotentId string) error {
    r._idempotentId = _idempotentId
    r.Set("idempotent_id", _idempotentId)
    return nil
}

// IdempotentId Getter
func (r AlibabaLatourStrategyIssueAPIRequest) GetIdempotentId() string {
    return r._idempotentId
}
// Channel Setter
// 发放渠道
func (r *AlibabaLatourStrategyIssueAPIRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r AlibabaLatourStrategyIssueAPIRequest) GetChannel() string {
    return r._channel
}
// UserId Setter
// 除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
func (r *AlibabaLatourStrategyIssueAPIRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaLatourStrategyIssueAPIRequest) GetUserId() string {
    return r._userId
}
// TransformedUserType Setter
// 转换用户类型
func (r *AlibabaLatourStrategyIssueAPIRequest) SetTransformedUserType(_transformedUserType string) error {
    r._transformedUserType = _transformedUserType
    r.Set("transformed_user_type", _transformedUserType)
    return nil
}

// TransformedUserType Getter
func (r AlibabaLatourStrategyIssueAPIRequest) GetTransformedUserType() string {
    return r._transformedUserType
}
// NeedIdentifyRisk Setter
// 是否需要过安全
func (r *AlibabaLatourStrategyIssueAPIRequest) SetNeedIdentifyRisk(_needIdentifyRisk bool) error {
    r._needIdentifyRisk = _needIdentifyRisk
    r.Set("need_identify_risk", _needIdentifyRisk)
    return nil
}

// NeedIdentifyRisk Getter
func (r AlibabaLatourStrategyIssueAPIRequest) GetNeedIdentifyRisk() bool {
    return r._needIdentifyRisk
}
// UserNick Setter
// 除非有特殊申请，默认不允许使用该参数，请通过用户授权token传递用户信息
func (r *AlibabaLatourStrategyIssueAPIRequest) SetUserNick(_userNick string) error {
    r._userNick = _userNick
    r.Set("user_nick", _userNick)
    return nil
}

// UserNick Getter
func (r AlibabaLatourStrategyIssueAPIRequest) GetUserNick() string {
    return r._userNick
}
// StrategyCode Setter
// 投放计划code
func (r *AlibabaLatourStrategyIssueAPIRequest) SetStrategyCode(_strategyCode string) error {
    r._strategyCode = _strategyCode
    r.Set("strategy_code", _strategyCode)
    return nil
}

// StrategyCode Getter
func (r AlibabaLatourStrategyIssueAPIRequest) GetStrategyCode() string {
    return r._strategyCode
}
// UserType Setter
// 用户类型
func (r *AlibabaLatourStrategyIssueAPIRequest) SetUserType(_userType string) error {
    r._userType = _userType
    r.Set("user_type", _userType)
    return nil
}

// UserType Getter
func (r AlibabaLatourStrategyIssueAPIRequest) GetUserType() string {
    return r._userType
}
// SelectedBenefitCode Setter
// 指定发放权益code
func (r *AlibabaLatourStrategyIssueAPIRequest) SetSelectedBenefitCode(_selectedBenefitCode string) error {
    r._selectedBenefitCode = _selectedBenefitCode
    r.Set("selected_benefit_code", _selectedBenefitCode)
    return nil
}

// SelectedBenefitCode Getter
func (r AlibabaLatourStrategyIssueAPIRequest) GetSelectedBenefitCode() string {
    return r._selectedBenefitCode
}