package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type Phoalb struct {
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

func (db *DB) QueryPhoalbByKeyAndUserId(key string, userid int) (note Phoalb, err error) {
	return note, db.db.Model(&Phoalb{}).Where("`key` = ? and user_id = ?", key, userid).Take(&note).Error
}

func (db *DB) QueryPhoalbByKey(key string) (note Phoalb, err error) {
	return note, db.db.Model(&Phoalb{}).Where("`key` = ? ", key).Take(&note).Error
}

func (db *DB) AllPhoalbVisitCount(key string) error {
	return db.db.Model(&Phoalb{}).Where("`key` = ?", key).UpdateColumn("visit", gorm.Expr("visit + 1")).Error
}

func (db *DB) DelPhoalbByKey(key string, userid int) (error) {
	return db.db.Delete(Phoalb{}, "`key` = ? and user_id = ? ", key, userid).Error
}
func (db *DB) QueryPhoalbsByPage(page, limit int, title string) (note []*Phoalb, err error) {
	return note, db.db.Model(&Phoalb{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Offset((page - 1) * limit).Limit(limit).Order("updated_at DESC").Find(&note).Error
}
func (db *DB) QueryPhoalbsByPageAndType(page, limit, visit int, title string) (note []*Phoalb, err error) {
	return note, db.db.Model(&Phoalb{}).Where("visit = ?", visit).Offset((page - 1) * limit).Limit(limit).Order("updated_at DESC").Find(&note).Error
}
func (db *DB) QueryPhoalbs() (note []*Phoalb, err error) {
	return note, db.db.Model(&Phoalb{}).Where("1=1").Order("updated_at DESC").Find(&note).Error
}
func (db *DB) QueryPhoalbsCount(title string) (cnt int, err error) {
	return cnt, db.db.Model(&Phoalb{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Offset(-1).Limit(-1).Count(&cnt).Error
}
func (db *DB) QueryPhoalbsCountByType(visit int) (cnt int, err error) {
	return cnt, db.db.Model(&Phoalb{}).Where("visit = ?", visit).Offset(-1).Limit(-1).Count(&cnt).Error
}

func (db *DB) SavePhoalb(n *Phoalb) error {
	return db.db.Save(n).Error
}
