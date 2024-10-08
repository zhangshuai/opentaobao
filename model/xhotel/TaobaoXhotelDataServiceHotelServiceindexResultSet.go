package xhotel

import (
	"sync"
)

// TaobaoXhotelDataServiceHotelServiceindexResultSet 结构体
type TaobaoXhotelDataServiceHotelServiceindexResultSet struct {
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// firstResult
	FirstResult *TopAdsHtlDataQueryResult `json:"first_result,omitempty" xml:"first_result,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoXhotelDataServiceHotelServiceindexResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelDataServiceHotelServiceindexResultSet)
	},
}

// GetTaobaoXhotelDataServiceHotelServiceindexResultSet() 从对象池中获取TaobaoXhotelDataServiceHotelServiceindexResultSet
func GetTaobaoXhotelDataServiceHotelServiceindexResultSet() *TaobaoXhotelDataServiceHotelServiceindexResultSet {
	return poolTaobaoXhotelDataServiceHotelServiceindexResultSet.Get().(*TaobaoXhotelDataServiceHotelServiceindexResultSet)
}

// ReleaseTaobaoXhotelDataServiceHotelServiceindexResultSet 释放TaobaoXhotelDataServiceHotelServiceindexResultSet
func ReleaseTaobaoXhotelDataServiceHotelServiceindexResultSet(v *TaobaoXhotelDataServiceHotelServiceindexResultSet) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.FirstResult = nil
	v.Success = false
	poolTaobaoXhotelDataServiceHotelServiceindexResultSet.Put(v)
}
