package cmd

import (
	"fmt"
	"log"

	"github.com/22Fariz22/mycloud/cli/mycloud/pkg"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list of all files on the cloud",
	Long:  `no arguments needed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")

		list, err := pkg.List()
		if err != nil {
			log.Fatal("error in list.go: ", err)
			fmt.Println("произошла ошибка")
		}

		for el := range list {
			fmt.Println(el)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
