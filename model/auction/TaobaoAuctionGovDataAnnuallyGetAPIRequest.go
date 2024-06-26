package auction

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionGovDataAnnuallyGetAPIRequest 按年统计法院拍卖数据 API请求
// taobao.auction.gov.data.annually.get
//
// 按月统计法院拍卖数据 包含：
// 标的件数统计：发布标的件数、结束标的件数、开拍标的件数
// 竞价实况：预计成交金额、出价次数、报名人数
// 在线标的：在线标的件数、意向用户数、网拍围观人次
//
// 最长6年，年起始时间2017年
type TaobaoAuctionGovDataAnnuallyGetAPIRequest struct {
	model.Params
	// 法院名称
	_courtName string
	// 开始年份
	_startYear string
	// 结束年份
	_endYear string
	// 统计数据是够包含下属法院
	_isIncludeSub bool
}

// NewTaobaoAuctionGovDataAnnuallyGetRequest 初始化TaobaoAuctionGovDataAnnuallyGetAPIRequest对象
func NewTaobaoAuctionGovDataAnnuallyGetRequest() *TaobaoAuctionGovDataAnnuallyGetAPIRequest {
	return &TaobaoAuctionGovDataAnnuallyGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAuctionGovDataAnnuallyGetAPIRequest) Reset() {
	r._courtName = ""
	r._startYear = ""
	r._endYear = ""
	r._isIncludeSub = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAuctionGovDataAnnuallyGetAPIRequest) GetApiMethodName() string {
	return "taobao.auction.gov.data.annually.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAuctionGovDataAnnuallyGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAuctionGovDataAnnuallyGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCourtName is CourtName Setter
// 法院名称
func (r *TaobaoAuctionGovDataAnnuallyGetAPIRequest) SetCourtName(_courtName string) error {
	r._courtName = _courtName
	r.Set("court_name", _courtName)
	return nil
}

// GetCourtName CourtName Getter
func (r TaobaoAuctionGovDataAnnuallyGetAPIRequest) GetCourtName() string {
	return r._courtName
}

// SetStartYear is StartYear Setter
// 开始年份
func (r *TaobaoAuctionGovDataAnnuallyGetAPIRequest) SetStartYear(_startYear string) error {
	r._startYear = _startYear
	r.Set("start_year", _startYear)
	return nil
}

// GetStartYear StartYear Getter
func (r TaobaoAuctionGovDataAnnuallyGetAPIRequest) GetStartYear() string {
	return r._startYear
}

// SetEndYear is EndYear Setter
// 结束年份
func (r *TaobaoAuctionGovDataAnnuallyGetAPIRequest) SetEndYear(_endYear string) error {
	r._endYear = _endYear
	r.Set("end_year", _endYear)
	return nil
}

// GetEndYear EndYear Getter
func (r TaobaoAuctionGovDataAnnuallyGetAPIRequest) GetEndYear() string {
	return r._endYear
}

// SetIsIncludeSub is IsIncludeSub Setter
// 统计数据是够包含下属法院
func (r *TaobaoAuctionGovDataAnnuallyGetAPIRequest) SetIsIncludeSub(_isIncludeSub bool) error {
	r._isIncludeSub = _isIncludeSub
	r.Set("is_include_sub", _isIncludeSub)
	return nil
}

// GetIsIncludeSub IsIncludeSub Getter
func (r TaobaoAuctionGovDataAnnuallyGetAPIRequest) GetIsIncludeSub() bool {
	return r._isIncludeSub
}

var poolTaobaoAuctionGovDataAnnuallyGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAuctionGovDataAnnuallyGetRequest()
	},
}

// GetTaobaoAuctionGovDataAnnuallyGetRequest 从 sync.Pool 获取 TaobaoAuctionGovDataAnnuallyGetAPIRequest
func GetTaobaoAuctionGovDataAnnuallyGetAPIRequest() *TaobaoAuctionGovDataAnnuallyGetAPIRequest {
	return poolTaobaoAuctionGovDataAnnuallyGetAPIRequest.Get().(*TaobaoAuctionGovDataAnnuallyGetAPIRequest)
}

// ReleaseTaobaoAuctionGovDataAnnuallyGetAPIRequest 将 TaobaoAuctionGovDataAnnuallyGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAuctionGovDataAnnuallyGetAPIRequest(v *TaobaoAuctionGovDataAnnuallyGetAPIRequest) {
	v.Reset()
	poolTaobaoAuctionGovDataAnnuallyGetAPIRequest.Put(v)
}
