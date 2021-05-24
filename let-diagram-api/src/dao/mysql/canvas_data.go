package mysql

import (
    `github.com/jinzhu/gorm`
    `lets_diagram/src/models`
)

// InsertNewCanvasData 向 canvas_data 表中插入一条新纪录
func InsertNewCanvasData(db *gorm.DB) (*models.CanvasData, error) {
    cd := &models.CanvasData{
        Data: "{}",
        Version: 0,
    }
    err := db.Create(cd).Error
    return cd, err
}

// SelectCanvasDataByID 根据 id 获取 data
func SelectCanvasDataByID(db *gorm.DB, id uint) (string, error) {
    cd := &models.CanvasData{}
    err := db.Select("data").Where("id = ?", id).First(cd).Error
    return cd.Data, err
}

// UpdateCanvasDataByID 更新 canvas_data 表中的某条记录
func UpdateCanvasDataByID(db *gorm.DB, id uint, newData string, newVersion int64) error {
    return db.Model(&models.CanvasData{}).Where("id = ?", id).Update(map[string]interface{}{
        "data": newData,
        "version": newVersion,
    }).Error
}

// SelectCanvasData 按照主键查询
func SelectCanvasData(db *gorm.DB, canvasID uint) (*models.CanvasData, error) {
    canvas := &models.CanvasData{}
    err := db.Where("id = ?", canvasID).First(canvas).Error
    return canvas, err
}

// DeleteCanvasData 删除特定记录
func DeleteCanvasData(db *gorm.DB, id uint) error {
    err := db.Where("id = ?", id).Delete(&models.CanvasData{}).Error
    return err
}
