package cmd

import (
	"fmt"
	"log"

	"github.com/22Fariz22/mycloud/cli/pkg"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "del",
	Long:  `as an argument, specify the full path to the file on the cloud`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")

		_, err := pkg.Delete(args)
		if err != nil {
			log.Fatal("err in download.go : ", err)
			fmt.Println("произошла ошибка")
			return
		}

		fmt.Println("Файл успешно удален")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
