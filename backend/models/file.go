package models

import "time"

type File struct {
	ID       uint64    `gorm:"column:id; primaryKey"`
	MemberID uint64    `gorm:"column:member_id"`
	Name     string    `gorm:"column:name"`
	Value    string    `gorm:"-"`
	Path     string    `gorm:"column:path"`
	Date     time.Time `gorm:"column:creation_date"`
}
