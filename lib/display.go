package lib

import (
	"strings"

	"github.com/acarl005/textcol"
	"github.com/diegomagdaleno/tagger/converters"
)

func getBaseName(fullPath string) string {
	splitPath := strings.Split(fullPath, "/")
	return splitPath[len(splitPath)-1]
}

func fileString(file converters.FileProperties) string {
	// We color it with the first color of the tag
	coloredKey := strings.ToLower(file.Tags[0])

	colors := ColorOfTag[coloredKey]

	baseName := getBaseName(file.Name)

	displayString := []string{colors[0], "● ", baseName, Reset}

	return strings.Join(displayString, "")

}

func formatedElements(elements []converters.FileProperties) []string {
	var formattedStrings []string

	for i := range elements {
		stringToAppend := fileString(elements[i])
		formattedStrings = append(formattedStrings, stringToAppend)
	}
	return formattedStrings
}

func InitialDisplay(allElements []converters.FileProperties) {
	itemsTarget := formatedElements(allElements)
	textcol.PrintColumns(&itemsTarget, 5)
}
