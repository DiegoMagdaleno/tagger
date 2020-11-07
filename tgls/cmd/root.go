package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/diegomagdaleno/tagger/converters"

	"github.com/diegomagdaleno/tagger/tgls/lib"

	"github.com/spf13/cobra"
)

var targetTag string
var exclusive bool
var noPretty bool

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

		_, err := os.Stat(targetDirectory)
		if err != nil {
			if strings.Contains(err.Error(), "no such file or directory") {
				fmt.Fprintln(os.Stderr, "No such file or directory")
				os.Exit(1)
			}
		}

		if targetTag == "all" && exclusive {
			fmt.Fprintln(os.Stderr, "Can't specify exclussivness without specifing a tag to search for exclusively")
			os.Exit(1)
		}

		initialFileList := converters.GetFinalArrayOfFiles(targetDirectory)

		switch {
		case targetTag != "all" && !exclusive:
			fileList = lib.SearchForFilesWithTags(initialFileList, targetTag)
		case targetTag != "all" && exclusive:
			fileList = lib.SearchForFilesWithTagsExclusively(initialFileList, targetTag)
		default:
			fileList = initialFileList
		}

		if noPretty {
			lib.RawDisplay(fileList)
		} else {
			lib.InitialDisplay(fileList)
		}

	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&targetTag, "search", "s", "all", "Allows you to search for files with a specific tag")
	rootCmd.Flags().BoolVarP(&exclusive, "exclusive", "e", false, "Only show files that contain a specific tag exclusively")
	rootCmd.Flags().BoolVarP(&noPretty, "no-pretty", "n", false, "Disables file pretty printing and prints absolue paths")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
