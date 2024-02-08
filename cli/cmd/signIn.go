/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/22Fariz22/mycloud/cli/pkg"
	pb "github.com/22Fariz22/mycloud/proto"

	"github.com/spf13/cobra"
)

var login string
var password string

// signInCmd represents the signIn command
var signInCmd = &cobra.Command{
	Use:   "signIn",
	Short: "log in to your account",
	Long:  `provide two arguments: email and password`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("signIn called")

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
