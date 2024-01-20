package cmd

import (
	"fmt"
	"log"

	"github.com/22Fariz22/mycloud/cli/mycloud/pkg"
	"github.com/spf13/cobra"
)

// logoutCmd exit session
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "exit session",
	Long:  `no arguments needed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logout called")

		_,err:=pkg.Logout()
		if err!=nil{
			log.Fatal("err in logout.go : ", err)
			fmt.Println("произошла ошибка")
		}

		fmt.Println("Вы успешно вышли из сессии")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
