package pkg

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	pb "github.com/22Fariz22/mycloud/cli/proto"
	"google.golang.org/grpc/metadata"
)

// Добавить интерфейс? в интерфейса можно ьыло бы добавить вызов?

func SignUp(c pb.UserServiceClient, input *pb.RegisterRequest) (string, error) {
	_, err := c.Register(context.Background(), input)
	if err != nil {
		fmt.Println("failed to register")
		return "", err
	}
	fmt.Println("register is successful")
	return "", nil
}

func SignIn(c pb.UserServiceClient, input *pb.LoginRequest) error {
	resp, err := c.Login(context.Background(), input)
	if err != nil {
		fmt.Println("failed to login")
		return err
	}

	//пишем в файл session_id
	f, err := os.Create("session.txt")
	if err != nil {
		fmt.Println("failed to login")
		return err
	}
	_, err = f.WriteString(resp.SessionId)
	if err != nil {
		err = errors.New("возможно вы еще не зарегистрировались")
		log.Println(err)
		return err
	}

	fmt.Println("login successfully")
	return nil
}

// Logout выход из session
func Logout(c pb.UserServiceClient, input *pb.LogoutRequest) error {
	ctx, err := GetSessionAndPutInMD()
	if err != nil {
		fmt.Println("failed to logout")
		return err
	}

	_, err = c.Logout(ctx, &pb.LogoutRequest{})
	if err != nil {
		fmt.Println("failed to logout")
		return err
	}
	fmt.Println("logout is successful")
	return nil
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

// GetSessionAndPutInMD читает файл session_id, ищет сессию ,вставляет session_id в метаданные и возвращает context
func GetSessionAndPutInMD() (context.Context, error) {
	data, err := os.ReadFile("session.txt")
	if err != nil {
		log.Println("err in ioutil.ReadFile:", err)
		return nil, err
	}

	//вставляем наш session_id в metadata
	md := metadata.New(map[string]string{"session_id": string(data)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	return ctx, nil
}
