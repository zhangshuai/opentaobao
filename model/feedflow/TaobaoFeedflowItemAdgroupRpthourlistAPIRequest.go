package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadgrouprpthourlistAPIRequest 超级推荐【商品推广】单元分时报表查询 API请求
// taobao.feedflow.item.adgroup.rpthourlist
//
// 广告主推广组分时数据查询，支持广告主查询最近90天内某一天的单元维度分时报表数据
type TaobaofeedflowitemadgrouprpthourlistAPIRequest struct {
	model.Params
	// 查询参数
	_rptQuery *RptQueryDto
}

// NewTaobaofeedflowitemadgrouprpthourlistRequest 初始化TaobaofeedflowitemadgrouprpthourlistAPIRequest对象
func NewTaobaofeedflowitemadgrouprpthourlistRequest() *TaobaofeedflowitemadgrouprpthourlistAPIRequest {
	return &TaobaofeedflowitemadgrouprpthourlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemadgrouprpthourlistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.rpthourlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemadgrouprpthourlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemadgrouprpthourlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQuery is RptQuery Setter
// 查询参数
func (r *TaobaofeedflowitemadgrouprpthourlistAPIRequest) SetRptQuery(_rptQuery *RptQueryDto) error {
	r._rptQuery = _rptQuery
	r.Set("rpt_query", _rptQuery)
	return nil
}

// GetRptQuery RptQuery Getter
func (r TaobaofeedflowitemadgrouprpthourlistAPIRequest) GetRptQuery() *RptQueryDto {
	return r._rptQuery
}
