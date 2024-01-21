package cmd

import (
	"fmt"

	"github.com/22Fariz22/mycloud/cli/pkg"
	"github.com/spf13/cobra"
	"log"
)

// shareCmd represents the share command
var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "share your file on the cloud",
	Long:  `in argument indicate the email with whom you want to share the file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("share called")

		msg, err := pkg.Share(args)
		if err != nil {
			log.Fatal("err in share.go : ", err)
			fmt.Println("произошла ошибка")
			return
		}

		fmt.Println("ссылка на скачивание: ", msg)
	},
}

func init() {
	rootCmd.AddCommand(shareCmd)
}
