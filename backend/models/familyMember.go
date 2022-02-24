package models

type FamilyMember struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"column:name"`
	AccountID uint64 `gorm:"column:account_id"`
}
