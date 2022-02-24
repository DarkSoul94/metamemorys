package postgres

import (
	"errors"
	"io"
	"os"

	"github.com/DarkSoul94/metamemorys_backend/metamemory"
	"github.com/DarkSoul94/metamemorys_backend/models"
	"gorm.io/gorm"
)

type metaRepo struct {
	db *gorm.DB
}

func NewMetaRepository(db *gorm.DB) metamemory.MetaRepository {
	return &metaRepo{
		db: db,
	}
}

func (r *metaRepo) GetMemberList(accountID uint64) ([]*models.FamilyMember, error) {
	var members []*models.FamilyMember
	tx := r.db.Where(&models.FamilyMember{AccountID: accountID}).Find(&members)
	return members, tx.Error
}

func (r *metaRepo) CreateMember(member *models.FamilyMember) (uint64, error) {
	tx := r.db.Create(&member)
	return member.ID, tx.Error
}

func (r *metaRepo) CreateFile(file *models.File) error {
	osFile, err := os.Create(file.Path)
	defer osFile.Close()
	if err != nil {
		return errors.New("Failed create file " + file.Name)
	}

	osFile.Write([]byte(file.Value))

	tx := r.db.Create(&file)
	return tx.Error
}

func (r *metaRepo) GetMemberFiles(memberID uint64) ([]*models.File, error) {
	var memberFiles []*models.File
	tx := r.db.Where("member_id = ?", memberID).Find(&memberFiles)

	for _, file := range memberFiles {
		osFile, err := os.Open(file.Path)
		if err != nil {
			return nil, err
		}
		defer osFile.Close()

		packet := make([]byte, 10000)
		fileData := make([]byte, 0)

		for {
			byteCount, err := osFile.Read(packet)
			fileData = append(fileData, packet[:byteCount]...)
			if err == io.EOF { // если конец файла
				break // выходим из цикла
			}
		}

		file.Value = string(fileData)
	}

	return memberFiles, tx.Error
}
