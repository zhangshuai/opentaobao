package qimen

// WaveNumReportRequest 
type WaveNumReportRequest struct {
    // 波次号
    WaveNum   string `json:"waveNum,omitempty" xml:"waveNum,omitempty"`
    // 发货单号
    Orders   []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
    // 扩展属性
    ExtendProps   *TaobaoQimenWavenumReportMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
