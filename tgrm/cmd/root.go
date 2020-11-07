package cmd

import (
	"fmt"
	"os"

	"github.com/diegomagdaleno/tagger/converters"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tgrm",
	Short: "tgrm allows you to delete a tag (or tags) from a file",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("At least 2 arguments are required, but %v where providen", len(args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		tag := args[1]
		_, err := os.Stat(file)
		if err != nil {
			fmt.Printf("%v: %v: %v\n", os.Args[0], file, err)
			os.Exit(1)
		}
		err = converters.RemoveTag(tag, file)
		if err != nil {
			fmt.Printf("%v: %v: %v\n", os.Args[0], file, err)
			os.Exit(1)
		}
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
