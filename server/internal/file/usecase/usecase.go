package usecase

import "github.com/22Fariz22/mycloud/server/internal/file"

type fileUC struct {
	pg file.FileRepository
}

func NewFileUC(pg file.FileRepository) file.UseCase {
	return &fileUC{
		pg: pg,
	}
}

func (u *fileUC) GetList() {}

func (u *fileUC) Upload() {}

func (u *fileUC) Download() {}

func (u *fileUC) Share() {}

func (u *fileUC) Delete() {}
