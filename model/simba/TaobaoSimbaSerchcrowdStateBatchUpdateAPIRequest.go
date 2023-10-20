package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest 单品搜索人群修改状态 API请求
// taobao.simba.serchcrowd.state.batch.update
//
// 暂停或启用单品推广搜索人群溢价
type TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest struct {
	model.Params
	// 需要修改出价的人群包id,批量传入时用,分割
	_adgroupCrowdIds []string
	// 被操作者的淘宝昵称
	_nick string
	// 推广单元id
	_adgroupId int64
	// 人群状态,0:暂停;1:启用
	_state int64
}

// NewTaobaoSimbaSerchcrowdStateBatchUpdateRequest 初始化TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest对象
func NewTaobaoSimbaSerchcrowdStateBatchUpdateRequest() *TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest {
	return &TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest) Reset() {
	r._adgroupCrowdIds = r._adgroupCrowdIds[:0]
	r._nick = ""
	r._adgroupId = 0
	r._state = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.serchcrowd.state.batch.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupCrowdIds is AdgroupCrowdIds Setter
// 需要修改出价的人群包id,批量传入时用,分割
func (r *TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest) SetAdgroupCrowdIds(_adgroupCrowdIds []string) error {
	r._adgroupCrowdIds = _adgroupCrowdIds
	r.Set("adgroup_crowd_ids", _adgroupCrowdIds)
	return nil
}

// GetAdgroupCrowdIds AdgroupCrowdIds Getter
func (r TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest) GetAdgroupCrowdIds() []string {
	return r._adgroupCrowdIds
}

// SetNick is Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetState is State Setter
// 人群状态,0:暂停;1:启用
func (r *TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest) SetState(_state int64) error {
	r._state = _state
	r.Set("state", _state)
	return nil
}

// GetState State Getter
func (r TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest) GetState() int64 {
	return r._state
}

var poolTaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaSerchcrowdStateBatchUpdateRequest()
	},
}

// GetTaobaoSimbaSerchcrowdStateBatchUpdateRequest 从 sync.Pool 获取 TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest
func GetTaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest() *TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest {
	return poolTaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest.Get().(*TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest)
}

// ReleaseTaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest 将 TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest(v *TaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest) {
	v.Reset()
	poolTaobaoSimbaSerchcrowdStateBatchUpdateAPIRequest.Put(v)
}
