package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemupdatedelistingtmallAPIRequest taobao.item.update.delisting.tmall API请求
// taobao.item.update.delisting.tmall
//
// * 单个商品下架&lt;br/&gt;    * 输入的num_iid必须属于当前会话用户
type TaobaoitemupdatedelistingtmallAPIRequest struct {
	model.Params
	// 商品数字ID，该参数必须
	_numIid int64
}

// NewTaobaoitemupdatedelistingtmallRequest 初始化TaobaoitemupdatedelistingtmallAPIRequest对象
func NewTaobaoitemupdatedelistingtmallRequest() *TaobaoitemupdatedelistingtmallAPIRequest {
	return &TaobaoitemupdatedelistingtmallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitemupdatedelistingtmallAPIRequest) GetApiMethodName() string {
	return "taobao.item.update.delisting.tmall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitemupdatedelistingtmallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitemupdatedelistingtmallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNumIid is NumIid Setter
// 商品数字ID，该参数必须
func (r *TaobaoitemupdatedelistingtmallAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoitemupdatedelistingtmallAPIRequest) GetNumIid() int64 {
	return r._numIid
}
