package cmd

import (
	"fmt"
	"log"

	"github.com/22Fariz22/mycloud/cli/pkg"
	pb "github.com/22Fariz22/mycloud/proto"

	"github.com/spf13/cobra"
)

// logoutCmd exit session
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "exit session",
	Long:  `no arguments needed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logout called")

		c := pkg.ConnGRPCServer()

		err := pkg.Logout(c, &pb.LogoutRequest{})
		if err != nil {
			log.Fatal(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
