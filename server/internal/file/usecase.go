package file

type UseCase interface {
	GetList()
	Upload()
	Download()
	Share()
	Delete()
}
