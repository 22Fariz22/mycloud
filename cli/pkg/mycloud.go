package pkg

import "fmt"

// Добавить интерфейс? в интерфейса можно ьыло бы добавить вызов?

func SignUp(args []string) (string, error) {

	return "", nil
}

func SignIn(args []string) (string, error) {

	return "", nil
}

func Logout() (string, error) {

	return "", nil
}

func List() ([]string, error) {

	return []string{}, nil
}

func Download(args []string) (string, error) {
	//func Open(name string) (*File, error)

	if len(args) == 0 {
		fmt.Println("вы не указали путь к файлу")
		return "", nil
	}

	return "", nil
}

func Upload() (string, error) {

	return "", nil
}

func Delete(args []string) (string, error) {

	return "", nil
}

// Share возвращает ссылку на скачивание и ошибку
func Share(args []string) (string, error) {

	return "", nil
}
