package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchanneldistributorproductdistributeAPIRequest 分销商基于渠道产品铺货到商品 API请求
// alibaba.ascp.channel.distributor.product.distribute
//
// 分销商基于渠道产品铺货到商品
type AlibabaascpchanneldistributorproductdistributeAPIRequest struct {
	model.Params
	// 请求参数
	_itemDistributePublishRequest *ItemDistributePublishRequest
}

// NewAlibabaascpchanneldistributorproductdistributeRequest 初始化AlibabaascpchanneldistributorproductdistributeAPIRequest对象
func NewAlibabaascpchanneldistributorproductdistributeRequest() *AlibabaascpchanneldistributorproductdistributeAPIRequest {
	return &AlibabaascpchanneldistributorproductdistributeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchanneldistributorproductdistributeAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.distributor.product.distribute"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchanneldistributorproductdistributeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchanneldistributorproductdistributeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemDistributePublishRequest is ItemDistributePublishRequest Setter
// 请求参数
func (r *AlibabaascpchanneldistributorproductdistributeAPIRequest) SetItemDistributePublishRequest(_itemDistributePublishRequest *ItemDistributePublishRequest) error {
	r._itemDistributePublishRequest = _itemDistributePublishRequest
	r.Set("item_distribute_publish_request", _itemDistributePublishRequest)
	return nil
}

// GetItemDistributePublishRequest ItemDistributePublishRequest Getter
func (r AlibabaascpchanneldistributorproductdistributeAPIRequest) GetItemDistributePublishRequest() *ItemDistributePublishRequest {
	return r._itemDistributePublishRequest
}