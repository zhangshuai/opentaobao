package traderate

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索评价信息 API请求
taobao.traderates.get

搜索评价信息，只能获取距今180天内的评价记录(只支持查询卖家给出或得到的评价)。
*/
type TaobaoTraderatesGetRequest struct {
    model.Params
    // 需返回的字段列表。可选值：TradeRate 结构中的所有字段，多个字段之间用“,”分隔
    fields   []string
    // 评价类型。可选值:get(得到),give(给出)
    rateType   string
    // 评价者角色即评价的发起方。可选值:seller(卖家),buyer(买家)。 当 give buyer 以买家身份给卖家的评价； 当 get seller 以买家身份得到卖家给的评价； 当 give seller 以卖家身份给买家的评价； 当 get buyer 以卖家身份得到买家给的评价。
    role   string
    // 评价结果。可选值:good(好评),neutral(中评),bad(差评)
    result   string
    // 页码。取值范围:大于零的整数最大限制为200; 默认值:1
    pageNo   int64
    // 每页获取条数。默认值40，最小值1，最大值150。
    pageSize   int64
    // 评价开始时。如果只输入开始时间，那么能返回开始时间之后的评价数据。
    startDate   string
    // 评价结束时间。如果只输入结束时间，那么全部返回所有评价数据。
    endDate   string
    // 交易订单id，可以是父订单id号，也可以是子订单id号
    tid   int64
    // 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取评价信息，效率在原有的基础上有80%的提升。
    useHasNext   bool
    // 商品的数字ID
    numIid   int64
}

// 初始化TaobaoTraderatesGetRequest对象
func NewTaobaoTraderatesGetRequest() *TaobaoTraderatesGetRequest{
    return &TaobaoTraderatesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTraderatesGetRequest) GetApiMethodName() string {
    return "taobao.traderates.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTraderatesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需返回的字段列表。可选值：TradeRate 结构中的所有字段，多个字段之间用“,”分隔
func (r *TaobaoTraderatesGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoTraderatesGetRequest) GetFields() []string {
    return r.fields
}
// RateType Setter
// 评价类型。可选值:get(得到),give(给出)
func (r *TaobaoTraderatesGetRequest) SetRateType(rateType string) error {
    r.rateType = rateType
    r.Set("rate_type", rateType)
    return nil
}

// RateType Getter
func (r TaobaoTraderatesGetRequest) GetRateType() string {
    return r.rateType
}
// Role Setter
// 评价者角色即评价的发起方。可选值:seller(卖家),buyer(买家)。 当 give buyer 以买家身份给卖家的评价； 当 get seller 以买家身份得到卖家给的评价； 当 give seller 以卖家身份给买家的评价； 当 get buyer 以卖家身份得到买家给的评价。
func (r *TaobaoTraderatesGetRequest) SetRole(role string) error {
    r.role = role
    r.Set("role", role)
    return nil
}

// Role Getter
func (r TaobaoTraderatesGetRequest) GetRole() string {
    return r.role
}
// Result Setter
// 评价结果。可选值:good(好评),neutral(中评),bad(差评)
func (r *TaobaoTraderatesGetRequest) SetResult(result string) error {
    r.result = result
    r.Set("result", result)
    return nil
}

// Result Getter
func (r TaobaoTraderatesGetRequest) GetResult() string {
    return r.result
}
// PageNo Setter
// 页码。取值范围:大于零的整数最大限制为200; 默认值:1
func (r *TaobaoTraderatesGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoTraderatesGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页获取条数。默认值40，最小值1，最大值150。
func (r *TaobaoTraderatesGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTraderatesGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// StartDate Setter
// 评价开始时。如果只输入开始时间，那么能返回开始时间之后的评价数据。
func (r *TaobaoTraderatesGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoTraderatesGetRequest) GetStartDate() string {
    return r.startDate
}
// EndDate Setter
// 评价结束时间。如果只输入结束时间，那么全部返回所有评价数据。
func (r *TaobaoTraderatesGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoTraderatesGetRequest) GetEndDate() string {
    return r.endDate
}
// Tid Setter
// 交易订单id，可以是父订单id号，也可以是子订单id号
func (r *TaobaoTraderatesGetRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoTraderatesGetRequest) GetTid() int64 {
    return r.tid
}
// UseHasNext Setter
// 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取评价信息，效率在原有的基础上有80%的提升。
func (r *TaobaoTraderatesGetRequest) SetUseHasNext(useHasNext bool) error {
    r.useHasNext = useHasNext
    r.Set("use_has_next", useHasNext)
    return nil
}

// UseHasNext Getter
func (r TaobaoTraderatesGetRequest) GetUseHasNext() bool {
    return r.useHasNext
}
// NumIid Setter
// 商品的数字ID
func (r *TaobaoTraderatesGetRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

// NumIid Getter
func (r TaobaoTraderatesGetRequest) GetNumIid() int64 {
    return r.numIid
}
