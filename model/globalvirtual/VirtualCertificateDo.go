package globalvirtual

import (
	"sync"
)

// VirtualCertificateDo 结构体
type VirtualCertificateDo struct {
	// provide download file
	File string `json:"file,omitempty" xml:"file,omitempty"`
	// code info
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// track order id
	OrderTrackId string `json:"order_track_id,omitempty" xml:"order_track_id,omitempty"`
	// remark info
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// modified time
	GmtModified int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// code start effective time
	StartTime int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// primary key
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// code end effective time
	EndTime int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// create time
	GmtCreate int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// trade order id
	TradeOrderLineId int64 `json:"trade_order_line_id,omitempty" xml:"trade_order_line_id,omitempty"`
}

var poolVirtualCertificateDo = sync.Pool{
	New: func() any {
		return new(VirtualCertificateDo)
	},
}

// GetVirtualCertificateDo() 从对象池中获取VirtualCertificateDo
func GetVirtualCertificateDo() *VirtualCertificateDo {
	return poolVirtualCertificateDo.Get().(*VirtualCertificateDo)
}

// ReleaseVirtualCertificateDo 释放VirtualCertificateDo
func ReleaseVirtualCertificateDo(v *VirtualCertificateDo) {
	v.File = ""
	v.Code = ""
	v.OrderTrackId = ""
	v.Remark = ""
	v.GmtModified = 0
	v.StartTime = 0
	v.Id = 0
	v.EndTime = 0
	v.GmtCreate = 0
	v.TradeOrderLineId = 0
	poolVirtualCertificateDo.Put(v)
}
