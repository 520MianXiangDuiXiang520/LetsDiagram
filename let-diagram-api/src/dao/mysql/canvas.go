package mysql

import (
    "github.com/520MianXiangDuiXiang520/GoTools/dao"
    "github.com/jinzhu/gorm"
    "lets_diagram/src/models"
)

// InsertNewCanvasInfo 向 canvas 表中插入一条记录
func InsertNewCanvasInfo(db *gorm.DB, canvas *models.Canvas) (*models.Canvas, error) {
	err := db.Create(canvas).Error
	return canvas, err
}

// SelectCanvasInfoByUserID 查询所有与 user 相关的 canvas 记录
func SelectCanvasInfoByUserID(db *gorm.DB, userID uint, offset, limit int) ([]models.Canvas, error) {
	ids, err := SelectCanvasIDByUserID(db, userID, offset, limit)
	if err != nil {
		return nil, err
	}
	result := make([]models.Canvas, 0)
	err = db.Model(&models.Canvas{}).Where("id IN (?)", ids).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SelectCanvasDataIDByCanvasID 根据 id 找到 canvas 对应的 data_id
func SelectCanvasDataIDByCanvasID(canvasID uint) (uint, error) {
	canvas := &models.Canvas{}
	err := dao.GetDB().Select("data_id").Where("id = ?", canvasID).First(canvas).Error
	return canvas.DataID, err
}

func SelectCanvasInfoByID(db *gorm.DB, id uint) (c *models.Canvas, e error) {
    c = &models.Canvas{}
    e =  db.Where("id = ?", id).First(&c).Error
    return c, e
}

// SelectCoverIDByCanvasID 根据 id 找到 canvas 对应的 cover_id
func SelectCoverIDByCanvasID(db *gorm.DB, canvasID uint) (uint, error) {
    canvas := &models.Canvas{}
    err := db.Select("cover_id").Where("id = ?", canvasID).First(canvas).Error
    return canvas.CoverID, err
}

func DeleteCanvasInfo(db *gorm.DB, id uint) error {
    err := db.Where("id = ?", id).Delete(models.Canvas{}).Error
    return err
}
