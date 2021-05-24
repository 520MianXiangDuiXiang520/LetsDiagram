package dao

import (
    `github.com/520MianXiangDuiXiang520/GoTools/dao`
    `lets_diagram/src/dao/mysql`
    `log`
)

func GetCoverData(canvasID uint) (string, bool) {
    db := dao.GetDB()
    coverID, err := mysql.SelectCoverIDByCanvasID(db, canvasID)
    if err != nil {
        log.Printf("cannot find the cover of canvas(%d): %v\n", canvasID, err)
        return "", false
    }
    cover, err := mysql.SelectCoverByID(db, coverID)
    if err != nil {
        log.Printf("cannot find cover(%d): %v\n", coverID, err)
        return "", false
    }
    return cover.Data, true
}
