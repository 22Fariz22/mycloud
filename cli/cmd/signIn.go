package cmd

import (
	"fmt"

	"github.com/22Fariz22/mycloud/cli/pkg"
	pb "github.com/22Fariz22/cli/proto"

	"github.com/spf13/cobra"
)

var login string
var password string

// signInCmd represents the signIn command
var signInCmd = &cobra.Command{
	Use:   "signin",
	Short: "in",
	Long:  `log in to your account, provide two arguments: email and password`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("signin called")

		c := pkg.ConnGRPCServer()

		pkg.SignIn(c, &pb.LoginRequest{
			Login:    login,
			Password: password,
		})
	},
}

func init() {
	rootCmd.AddCommand(signInCmd)
}
