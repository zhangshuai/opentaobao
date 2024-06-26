package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentShowGetbyshowidAPIRequest 迎客松根据节目id获取节目元数据 API请求
// yunos.tvpubadmin.content.show.getbyshowid
//
// 迎客松根据节目id获取节目元数据
type YunosTvpubadminContentShowGetbyshowidAPIRequest struct {
	model.Params
	// 节目字符串id
	_showId string
}

// NewYunosTvpubadminContentShowGetbyshowidRequest 初始化YunosTvpubadminContentShowGetbyshowidAPIRequest对象
func NewYunosTvpubadminContentShowGetbyshowidRequest() *YunosTvpubadminContentShowGetbyshowidAPIRequest {
	return &YunosTvpubadminContentShowGetbyshowidAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentShowGetbyshowidAPIRequest) Reset() {
	r._showId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentShowGetbyshowidAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.show.getbyshowid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentShowGetbyshowidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentShowGetbyshowidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShowId is ShowId Setter
// 节目字符串id
func (r *YunosTvpubadminContentShowGetbyshowidAPIRequest) SetShowId(_showId string) error {
	r._showId = _showId
	r.Set("show_id", _showId)
	return nil
}

// GetShowId ShowId Getter
func (r YunosTvpubadminContentShowGetbyshowidAPIRequest) GetShowId() string {
	return r._showId
}

var poolYunosTvpubadminContentShowGetbyshowidAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentShowGetbyshowidRequest()
	},
}

// GetYunosTvpubadminContentShowGetbyshowidRequest 从 sync.Pool 获取 YunosTvpubadminContentShowGetbyshowidAPIRequest
func GetYunosTvpubadminContentShowGetbyshowidAPIRequest() *YunosTvpubadminContentShowGetbyshowidAPIRequest {
	return poolYunosTvpubadminContentShowGetbyshowidAPIRequest.Get().(*YunosTvpubadminContentShowGetbyshowidAPIRequest)
}

// ReleaseYunosTvpubadminContentShowGetbyshowidAPIRequest 将 YunosTvpubadminContentShowGetbyshowidAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentShowGetbyshowidAPIRequest(v *YunosTvpubadminContentShowGetbyshowidAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentShowGetbyshowidAPIRequest.Put(v)
}
