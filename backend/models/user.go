package models

type User struct {
	ID       uint64 `gorm:"primaryKey"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email; unique"`
	PassHash string `gorm:"column:pass_hash"`
}
