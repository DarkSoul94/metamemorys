package usecase

import (
	"fmt"
	"os"
	"time"

	"github.com/DarkSoul94/metamemorys_backend/metamemory"
	"github.com/DarkSoul94/metamemorys_backend/models"
	"github.com/spf13/viper"
)

// Usecase ...
type metaUC struct {
	repo metamemory.MetaRepository
}

// NewUsecase ...
func NewMetaUsecase(repo metamemory.MetaRepository) metamemory.MetaUsecase {
	return &metaUC{
		repo: repo,
	}
}

func (u *metaUC) GetMemberList(accountID uint64) ([]*models.FamilyMember, error) {
	return u.repo.GetMemberList(accountID)
}

func (u *metaUC) CreateMember(member *models.FamilyMember) (uint64, error) {
	return u.repo.CreateMember(member)
}

func (u *metaUC) CreateFiles(userID, memberID uint64, files []*models.File) error {
	defaultPath := viper.GetString("app.store.path")
	pathToFolder := u.buildPathToFolder(defaultPath, userID, memberID)

	for _, file := range files {
		file.MemberID = memberID
		file.Date = time.Now().Truncate(time.Second)
		file.Path = fmt.Sprintf("%s/%s", pathToFolder, file.Name)
		if err := u.repo.CreateFile(file); err != nil {
			return err
		}
	}

	return nil
}

func (u *metaUC) buildPathToFolder(defaultPath string, userID, memberID uint64) string {
	pathToMember := fmt.Sprintf("%s/%d/%d", defaultPath, userID, memberID)
	if _, err := os.Stat(pathToMember); os.IsNotExist(err) {
		pathToUser := fmt.Sprintf("%s/%d", defaultPath, userID)
		if _, err := os.Stat(pathToUser); os.IsExist(err) {
			os.Mkdir(pathToMember, 0777)
		} else {
			os.Mkdir(pathToUser, 0777)
			os.Mkdir(pathToMember, 0777)
		}
	}
	return pathToMember
}

func (u *metaUC) GetMemberFiles(memberID uint64) ([]*models.File, error) {
	return u.repo.GetMemberFiles(memberID)
}
