package usecase

type UseCase interface{
	SignUp(args []string) (string, error)
	SignIn(args []string) (string, error)
	Logout() (string, error)
	List() ([]string, error)
	Download(args []string) (string, error)
	Upload() (string, error)
	Delete(args []string) (string, error)
	Share(args []string) (string, error)
}