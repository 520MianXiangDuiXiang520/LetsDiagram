package mysql

import (
    `github.com/jinzhu/gorm`
    `lets_diagram/src/models`
)

func SelectCoverByID(db *gorm.DB, id uint) (*models.Covers, error) {
    res := &models.Covers{}
    err := db.Where("id = ?", id).First(res).Error
    return res, err
}

func InsertNewCover(db *gorm.DB) (*models.Covers, error) {
    c := &models.Covers{}
    err := db.Create(&c).Error
    return c, err
}

func UpdateCoverByID(db *gorm.DB, id uint, data string) error {
    return db.Model(&models.Covers{}).Where("id = ?", id).Update("data", data).Error
}

func DeleteCover(db *gorm.DB, id uint) error {
    err := db.Where("id = ?", id).Delete(&models.Covers{}).Error
    return err
}

