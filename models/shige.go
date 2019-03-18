package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type ShiGe struct {
	Model
	Key     string `gorm:"unique_index;not null;"`
	UserID  int
	User    User
	Title   string
	Summary string `gorm:"type:text"`
	Content string `gorm:"type:text"`
	Files   string `gorm:"type:text"`
	Visit   int    `gorm:"default:0"`
	Praise  int    `gorm:"default:0"`
}

func (db *DB) QueryShiGeByKeyAndUserId(key string, userid int) (shige ShiGe, err error) {
	return shige, db.db.Model(&ShiGe{}).Where("`key` = ? and user_id = ?", key, userid).Take(&shige).Error
}

func (db *DB) QueryShiGeByKey(key string) (shige ShiGe, err error) {
	return shige, db.db.Model(&ShiGe{}).Where("`key` = ? ", key).Take(&shige).Error
}

func (db *DB) AllVisitOnShiGeCount(key string) error {
	return db.db.Model(&ShiGe{}).Where("`key` = ?", key).UpdateColumn("visit", gorm.Expr("visit + 1")).Error
}

func (db *DB) DelShiGeByKey(key string, userid int) (error) {
	return db.db.Delete(ShiGe{}, "`key` = ? and user_id = ? ", key, userid).Error
}
func (db *DB) QueryShiGesByPage(page, limit int, title string) (shige []*ShiGe, err error) {
	return shige, db.db.Model(&ShiGe{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Offset((page - 1) * limit).Limit(limit).Order("updated_at DESC").Find(&shige).Error
}
func (db *DB) QueryShiGesCount(title string) (cnt int, err error) {
	return cnt, db.db.Model(&ShiGe{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Offset(-1).Limit(-1).Count(&cnt).Error
}

func (db *DB) SaveShiGe(n *ShiGe) error {
	return db.db.Save(n).Error
}
