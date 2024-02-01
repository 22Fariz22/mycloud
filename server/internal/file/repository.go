package file

type FileRepository interface {
	GetList()
	Upload()
	Download()
	Share()
	Delete()
}
