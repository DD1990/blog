package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type Photos struct {
	Model
	Key     string `gorm:"not null;"`
	UserID  int
	User    User
	Title   string
	Summary string `gorm:"type:text"`
	Content string `gorm:"type:text"`
	Files   string `gorm:"type:text"`
	Visit   int    `gorm:"default:0"`
	Praise  int    `gorm:"default:0"`
}

func (db *DB) QueryPhotosByKeyAndUserId(key string, userid int) (note Photos, err error) {
	return note, db.db.Model(&Photos{}).Where("`key` = ? and user_id = ?", key, userid).Take(&note).Error
}

func (db *DB) QueryPhotosByKey(key string) (note Photos, err error) {
	return note, db.db.Model(&Photos{}).Where("`key` = ? ", key).Take(&note).Error
}

func (db *DB) QueryPhotosAllByKey(key string) (note []*Photos, err error) {
	return note, db.db.Model(&Photos{}).Where("`key` = ? ", key).Find(&note).Error
}

func (db *DB) AllPhotosVisitCount(key string) error {
	return db.db.Model(&Photos{}).Where("`key` = ?", key).UpdateColumn("visit", gorm.Expr("visit + 1")).Error
}

func (db *DB) DelPhotosByKey(key string, userid int) (error) {
	return db.db.Delete(Photos{}, "`key` = ? and user_id = ? ", key, userid).Error
}
func (db *DB) QueryPhotosByPage(page, limit int, title string) (note []*Photos, err error) {
	return note, db.db.Model(&Photos{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Offset((page - 1) * limit).Limit(limit).Order("updated_at DESC").Find(&note).Error
}
func (db *DB) QueryPhotosCount(title string) (cnt int, err error) {
	return cnt, db.db.Model(&Photos{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Offset(-1).Limit(-1).Count(&cnt).Error
}

func (db *DB) SavePhotos(n *Photos) error {
	return db.db.Save(n).Error
}
