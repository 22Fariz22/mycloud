package pkg

import "fmt"

// Добавить интерфейс? в интерфейса можно ьыло бы добавить вызов?

func SignUp(args []string) (string, error) {

	return "", nil
}

func SignIn(args []string) (string, error) {

	return "", nil
}

func SignOut() (string, error) {

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
	// Assuming you're uploading the file from a client, e.g. a browser, with Content-Type: application/x-www-form-urlencoded you should use FormFile instead of r.Form.Get which returns a *multipart.File value that contains the content of the file the client sent and which you can use to write that content to disk with io.Copy or what not.

	return "", nil
}

func Delete(args []string) (string, error) {

	return "", nil
}

//Share возвращает ссылку на скачивание и ошибку
func Share(args []string) (string, error) {

	return "", nil
}
