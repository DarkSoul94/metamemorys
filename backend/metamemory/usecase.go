package metamemory

import "github.com/DarkSoul94/metamemorys_backend/models"

// Usecase ...
type MetaUsecase interface {
	GetMemberList(accountID uint64) ([]*models.FamilyMember, error)
	CreateMember(member *models.FamilyMember) (uint64, error)

	CreateFiles(userID, memberID uint64, files []*models.File) error
	GetMemberFiles(memberID uint64) ([]*models.File, error)
}
