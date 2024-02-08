package cmd

import (
	"fmt"
	"log"

	"github.com/22Fariz22/mycloud/cli/pkg"
	pb "github.com/22Fariz22/mycloud/proto"
	"github.com/spf13/cobra"
)

var registerReq pb.LoginRequest

// signUpCmd represents the signUp command
var signUpCmd = &cobra.Command{
	Use:   "signup",
	Long:  `registration new an account in Mycloud, provide two arguments: email and password`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("signUp called")

		c := pkg.ConnGRPCServer()

		_, err := pkg.SignUp(c, &pb.RegisterRequest{
			Login:    registerReq.Login,
			Password: registerReq.Password,
		})

		if err != nil {
			log.Fatal("err in signUp.go : ", err)
			fmt.Println("произошла ошибка")
			return
		}

		fmt.Println("Вы успешно зарегистрировались")
	},
}

func init() {
	rootCmd.AddCommand(signUpCmd)
}
