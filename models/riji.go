package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type Riji struct {
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

func (db *DB) QueryRijiByKeyAndUserId(key string, userid int) (note Riji, err error) {
	return note, db.db.Model(&Riji{}).Where("`key` = ? and user_id = ?", key, userid).Take(&note).Error
}

func (db *DB) QueryRijiByKey(key string) (note Riji, err error) {
	return note, db.db.Model(&Riji{}).Where("`key` = ? ", key).Take(&note).Error
}

func (db *DB) QueryRijiByDate(date string) (note Riji, err error) {
	return note, db.db.Model(&Riji{}).Where("`created_at` like ? ", date).Take(&note).Error
}

func (db *DB) AllRijiVisitCount(key string) error {
	return db.db.Model(&Riji{}).Where("`key` = ?", key).UpdateColumn("visit", gorm.Expr("visit + 1")).Error
}

func (db *DB) DelRijiByKey(key string, userid int) (error) {
	return db.db.Delete(Riji{}, "`key` = ? and user_id = ? ", key, userid).Error
}
func (db *DB) QueryRijisByPage(page, limit int, title string) (note []*Riji, err error) {
	return note, db.db.Model(&Riji{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Offset((page - 1) * limit).Limit(limit).Order("updated_at DESC").Find(&note).Error
}
func (db *DB) QueryRijisCount(title string) (cnt int, err error) {
	return cnt, db.db.Model(&Riji{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Offset(-1).Limit(-1).Count(&cnt).Error
}

func (db *DB) SaveRiji(n *Riji) error {
	return db.db.Save(n).Error
}
