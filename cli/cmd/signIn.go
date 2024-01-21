/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/22Fariz22/mycloud/cli/pkg"
	"github.com/spf13/cobra"
)

// signInCmd represents the signIn command
var signInCmd = &cobra.Command{
	Use:   "signIn",
	Short: "log in to your account",
	Long:  `provide two arguments: email and password`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("signIn called")

		_, err := pkg.SignIn(args)
		if err != nil {
			log.Fatal("error in signIn: ", err)
			fmt.Println("произошла ошибка")
			return
		}

		fmt.Println("Вы успешно вошли в аккаунт")
	},
}

func init() {
	rootCmd.AddCommand(signInCmd)
}
