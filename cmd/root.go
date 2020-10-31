package cmd

import (
	"fmt"
	"os"

	"github.com/diegomagdaleno/tagger/converters"

	"github.com/diegomagdaleno/tagger/lib"

	"github.com/spf13/cobra"
)

var targetTag string

var rootCmd = &cobra.Command{
	Use:   "tagger",
	Short: "Tagger allows you to  manage macOS tags from the terminal",
	Run: func(cmd *cobra.Command, args []string) {
		var fileList []converters.FileProperties
		targetDirectory := func(args []string) string {
			if len(args) < 1 {
				return lib.GetCurrentDirectory()
			}
			return args[0]
		}(args)
		initialFileList := converters.GetFinalArrayOfFiles(targetDirectory)
		if targetTag != "all" {
			fileList = lib.SearchForFilesWithTags(initialFileList, targetTag)
		} else {
			fileList = initialFileList
		}

		lib.InitialDisplay(fileList)
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&targetTag, "search", "s", "all", "Allows you to search for files with a specific tag")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
