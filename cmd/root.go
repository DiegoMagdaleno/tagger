package cmd

import (
	"fmt"
	"os"

	"github.com/diegomagdaleno/tagger/converters"

	"github.com/diegomagdaleno/tagger/lib"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tagger",
	Short: "Tagger allows you to  manage macOS tags from the terminal",
	Run: func(cmd *cobra.Command, args []string) {
		targetDirectory := func(args []string) string {
			if len(args) < 1 {
				return lib.GetCurrentDirectory()
			}
			return args[0]
		}(args)
		initialFileList := converters.GetFinalArrayOfFiles(targetDirectory)
		lib.InitialDisplay(initialFileList)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
