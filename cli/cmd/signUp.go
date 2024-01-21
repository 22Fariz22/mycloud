package cmd

import (
	"fmt"
	"log"

	"github.com/22Fariz22/mycloud/cli/pkg"
	"github.com/spf13/cobra"
)

// signUpCmd represents the signUp command
var signUpCmd = &cobra.Command{
	Use:   "signUp",
	Short: "registration new an account in Mycloud",
	Long:  `provide two arguments: email and password`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("signUp called")

		_, err := pkg.SignUp(args)
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
