package iotticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iotticket"
)

/* 
服务商制定维修费方案 
cainiao.iot.ticket.sp.maintain.vtwo.create

服务商制定维修费方案
*/
func CainiaoIotTicketSpMaintainVtwoCreate(clt *core.SDKClient, req *iotticket.CainiaoIotTicketSpMaintainVtwoCreateRequest, session string) (*iotticket.CainiaoIotTicketSpMaintainVtwoCreateAPIResponse, error) {
    var resp iotticket.CainiaoIotTicketSpMaintainVtwoCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}