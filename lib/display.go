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
	var coloredBulletArray []string

	for i := range file.Tags {
		tag := strings.ToLower(file.Tags[i])
		coloredBulletArray = append(coloredBulletArray, ColorOfTag[tag][0], "●")
	}

	joinedBullets := strings.Join(coloredBulletArray, "")

	baseName := getBaseName(file.Name)

	displayString := []string{joinedBullets, " ", Reset, baseName}

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
	textcol.PrintColumns(&itemsTarget, 2)
}
