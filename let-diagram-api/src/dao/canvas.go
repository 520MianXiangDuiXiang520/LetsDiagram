package dao

import (
	"fmt"
	"github.com/520MianXiangDuiXiang520/GoTools/dao"
	json_diff "github.com/520MianXiangDuiXiang520/json-diff"
	"github.com/pkg/errors"
	"lets_diagram/src/dao/mysql"
	"lets_diagram/src/models"
	"lets_diagram/src/models/nottable"
	"log"
)

func DeleteCanvas(canvasID, userID uint) bool {
    var err error
    tx := dao.GetDB().Begin()
    defer func() {
        if err != nil {
            log.Printf("%+v", err)
            tx.Rollback()
        }
        tx.Commit()
    }()
    deleted, err := mysql.SelectCanvasInfoByID(tx, canvasID)
    if err != nil {
        err = errors.Wrap(err, fmt.Sprintf("can not find canvas where id = %d", canvasID))
        return false
    }
    // 删除 canvas_user
    err = mysql.DeleteCanvasUser(tx, canvasID, userID)
    if err != nil {
        return false
    }
    // 删除 canvas_info
    err = mysql.DeleteCanvasInfo(tx, canvasID)
    if err != nil {
        return false
    }
    // 删除 canvas_data
    err = mysql.DeleteCanvasData(tx, deleted.DataID)
    if err != nil {
        return false
    }
    // 删除 canvas_cover
    err = mysql.DeleteCover(tx, deleted.CoverID)
    if err != nil {
        return false
    }
    return true
}

// NewCanvas 新建 canvas
func NewCanvas(name string, userID uint) (*models.Canvas, bool) {
	db := dao.GetDB()
	tx := db.Begin()
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
			log.Printf("[dao/mysql/canvas] ROLLBACK NewCanvas() got a error: %v", err)
		}
		tx.Commit()
	}()

	// 新增 canvas_data 记录
	cd, err := mysql.InsertNewCanvasData(tx)
	if err != nil {
		return nil, false
	}

	// 新增 canvas_cover 记录
	cover, err := mysql.InsertNewCover(tx)
	if err != nil {
		return nil, false
	}

	newCanvas := &models.Canvas{
		DataID:  cd.ID,
		CoverID: cover.ID,
		Name:    name,
	}

	// 新增 canvas 记录
	canvasInfo, err := mysql.InsertNewCanvasInfo(tx, newCanvas)
	if err != nil {
		return nil, false
	}

	// 新增 canvas_user 记录
	err = mysql.InsertNewCanvasUser(tx, canvasInfo.ID, userID)
	if err != nil {
		return nil, false
	}
	return canvasInfo, true
}

func GetUserAllCanvasTotal(userID uint) int {
    db := dao.GetDB()
    res, err := mysql.CountUserAllCanvas(db, userID)
    if err != nil {
        log.Printf("fail to call GetUserAllCanvasTotal(%d), got error: %v", err)
        return 0
    }
    return res
}

// GetUserAllCanvas 获取用户的所有 canvas
func GetUserAllCanvas(userID uint, offset, limit int) []nottable.FullCanvas {
	db := dao.GetDB()
	canvas, err := mysql.SelectCanvasInfoByUserID(db, userID, offset, limit)
	if err != nil {
		log.Printf("[dao/mysql/canvas] get a error when select canvas info by userID(%d): %v ", userID, err)
		return nil
	}
	fcs := make([]nottable.FullCanvas, len(canvas))
	for i, c := range canvas {
		// cover, err := mysql.SelectCoverByID(db, c.CoverID)
		// if err != nil {
		// 	log.Printf("fail to get cover by %d, error is %+v", c.CoverID, err)
		// }
		fcs[i] = nottable.FullCanvas{
			ID:     c.ID,
			Name:   c.Name,
			Author: userID,
			// Cover:  cover.Data,
		}
	}
	return fcs
}

// GetCanvasData 获取某个 canvas 的数据
func GetCanvasData(canvasID uint) (*models.CanvasData, error) {
	db := dao.GetDB()
	dataID, err := mysql.SelectCanvasDataIDByCanvasID(canvasID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can not find canvasData: %d", canvasID))
	}
	canvas, err := mysql.SelectCanvasData(db, dataID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can not find %d canvasData", dataID))
	}
	return canvas, nil
}

func UpdateCanvasData(canvasID uint, newData *json_diff.JsonNode, newVersion int64) error {
	db := dao.GetDB()
	tx := db.Begin()
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
			log.Printf("[dao/mysql/canvas] ROLLBACK UpdateCanvasData() got a error: %v", err)
		}
		tx.Commit()
	}()
	newJsonString, err := json_diff.Marshal(newData)
	if err != nil {
		log.Printf("UpdateCanvasData() do json_diff.Marshal: %v", err)
		return err
	}
	// 获取 canvas_data.id
	cdID, err := mysql.SelectCanvasDataIDByCanvasID(canvasID)
	if err != nil {
		log.Printf("can not find canvas's (%d) data", canvasID)
	}
	// 更新 canvas_data.data
	err = mysql.UpdateCanvasDataByID(tx, cdID, string(newJsonString), newVersion)
	if err != nil {
		log.Printf("UpdateCanvasData() got : %v", err)
		return err
	}
	return nil
}

func UpdateCanvasCover(canvasID uint, data string) error {
	db := dao.GetDB()
	coverID, err := mysql.SelectCoverIDByCanvasID(db, canvasID)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("fail to get cover by canvasID(%d)", canvasID))
	}
	if err := mysql.UpdateCoverByID(db, coverID, data); err != nil {
		return errors.Wrap(err, fmt.Sprintf("fail to update cover by ID(%d)", coverID))
	}
	return nil
}
