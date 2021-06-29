package seaking

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/seaking"
)

/* 
跳转任务发布成功商品ID回传 
alibaba.seaking.task.report

跳转任务发布成功商品ID回传
*/
func AlibabaSeakingTaskReport(clt *core.SDKClient, req *seaking.AlibabaSeakingTaskReportRequest, session string) (*seaking.AlibabaSeakingTaskReportAPIResponse, error) {
    var resp seaking.AlibabaSeakingTaskReportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}