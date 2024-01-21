package cmd

import (
	"fmt"
	"log"

	"github.com/22Fariz22/mycloud/cli/pkg"
	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload upur file to Mycloud",
	Long:  `in argument provide full path to file in your computer`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upload called")

		file, err := pkg.Upload()
		if err != nil {
			log.Fatal("err in upload.go : ", err)
			fmt.Println("произошла ошибка")
			return
		}

		fmt.Println("файл успешно скачан:")
		fmt.Println(file)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
