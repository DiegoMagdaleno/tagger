package lib

import (
	"strings"

	"github.com/acarl005/textcol"
)

func getLastItem(allPaths []string) []string {
	var relativePaths []string
	for i := range allPaths {
		splitedPath := strings.Split(allPaths[i], "/")
		relativePaths = append(relativePaths, "\u2B24 "+splitedPath[len(splitedPath)-1])
	}
	return relativePaths
}

func InitialDisplay(files []string) {
	itemsTarget := getLastItem(files)
	textcol.PrintColumns(&itemsTarget, 5)
}
