package cmd

import (
	"fmt"
	"log"

	"github.com/22Fariz22/mycloud/cli/mycloud/pkg"
	"github.com/spf13/cobra"
)

// downloadCmd download file from cloud
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download file from cloud",
	Long:  `in argument provide full path to file in Mycloud `,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("download called")

		_, err := pkg.Download(args)
		if err != nil {
			log.Fatal("err in download.go : ", err)
			fmt.Println("произошла ошибка")
			return
		}

		fmt.Println("файл успешно скачан")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
