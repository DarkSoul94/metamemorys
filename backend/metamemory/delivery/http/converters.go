package http

import "github.com/DarkSoul94/metamemorys_backend/models"

func (h *MetaHandler) toOutMember(member *models.FamilyMember) outMember {
	return outMember{
		ID:   member.ID,
		Name: member.Name,
	}
}

func (h *MetaHandler) toModelsMember(member newMember, user *models.User) *models.FamilyMember {
	return &models.FamilyMember{
		Name:      member.Name,
		AccountID: user.ID,
	}
}

func (h *MetaHandler) toModelFiles(inpFiles inpFiles) []*models.File {
	var mFiles []*models.File = make([]*models.File, 0)
	for _, inpFile := range inpFiles.Files {
		mFiles = append(mFiles, &models.File{
			Name:  inpFile.Name,
			Value: inpFile.Value,
		})
	}

	return mFiles
}

func (h *MetaHandler) toOutFile(mFile *models.File) outFile {
	return outFile{
		ID:    mFile.ID,
		Name:  mFile.Name,
		Value: mFile.Value,
		Date:  mFile.Date,
	}
}
