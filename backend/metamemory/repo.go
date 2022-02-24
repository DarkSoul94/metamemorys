package metamemory

import "github.com/DarkSoul94/metamemorys_backend/models"

// Repository ...
type MetaRepository interface {
	GetMemberList(accountID uint64) ([]*models.FamilyMember, error)
	CreateMember(member *models.FamilyMember) (uint64, error)

	CreateFile(file *models.File) error
	GetMemberFiles(memberID uint64) ([]*models.File, error)
}
