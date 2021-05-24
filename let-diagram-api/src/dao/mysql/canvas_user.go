package mysql

import (
    `github.com/520MianXiangDuiXiang520/GoTools/dao`
    `github.com/jinzhu/gorm`
    `lets_diagram/src/models`
)

// CountUserAllCanvas 返回 user 有多少 canvas
func CountUserAllCanvas(db *gorm.DB, userID uint) (res int, err error) {
    err = db.Model(&models.UserCanvas{}).Where("user_id = ?", userID).Count(&res).Error
    return
}

// InsertNewCanvasUser 向 user_canvas 表中插入一条记录
func InsertNewCanvasUser(db *gorm.DB, canvasID, userID uint) error {
    return db.Create(&models.UserCanvas{
        UserID:   userID,
        CanvasID: canvasID,
    }).Error
}

// SelectByUserIDCanvasID SELECT * FROM user_canvas WHERE user_id = ? AND canvas_id = ?
func SelectByUserIDCanvasID(canvasID, userID uint) (*models.UserCanvas, error) {
    res := new(models.UserCanvas)
    err := dao.GetDB().Select("id").Where("canvas_id = ? AND user_id = ?", canvasID, userID).First(&res).Error
    return res, err
}

// SelectCanvasIDByUserID SELECT canvas_id FROM user_canvas WHERE user_id = ?
func SelectCanvasIDByUserID(db *gorm.DB, userID uint, offset, limit int) ([]uint, error) {
    canvasIDs := make([]*models.UserCanvas, 0)
    err := db.Select("canvas_id").Where("user_id = ?", userID).
        Offset(offset).Limit(limit).Find(&canvasIDs).Error
    res := make([]uint, len(canvasIDs))
    for i, id := range canvasIDs {
        res[i] = id.CanvasID
    }
    return res, err
}

// DeleteCanvasUser 删除一条记录
func DeleteCanvasUser(db *gorm.DB, canvasID, userID uint) error {
    err := db.Where("canvas_id = ? AND user_id = ?", canvasID, userID).Delete(&models.UserCanvas{}).Error
    return err
}