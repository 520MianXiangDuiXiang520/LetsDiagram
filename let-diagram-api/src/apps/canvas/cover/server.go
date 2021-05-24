package cover

import (
    `github.com/520MianXiangDuiXiang520/ginUtils`
    `lets_diagram/src/dao`
    `net/http`
)

func UpdateLogic(request *ginUtils.Request, response *ginUtils.Response) error {
    req := request.Req.(*UpdateRequestFields)
    err := dao.UpdateCanvasCover(req.CanvasID, req.Data)
    if err != nil {
        response.RespCode = http.StatusBadRequest
        response.Resp = UpdateResponseFields{
            Header: ginUtils.ParamErrorRespHeader,
        }
        return nil
    }
    response.RespCode = http.StatusOK
    response.Resp = UpdateResponseFields{
        Header: ginUtils.SuccessRespHeader,
    }
    return nil
}

func GetLogic(request *ginUtils.Request, response *ginUtils.Response) error {
    req := request.Req.(*GetRequestFields)
    cover, ok := dao.GetCoverData(req.CanvasID)
    if !ok {
        response.RespCode = http.StatusBadRequest
        response.Resp = GetResponseFields{Header: ginUtils.ParamErrorRespHeader, Cover: ""}
        return nil
    }
    response.RespCode = http.StatusOK
    response.Resp = GetResponseFields{Header: ginUtils.SuccessRespHeader, Cover: cover}
    return nil
}
