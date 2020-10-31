package lib

import (
	"reflect"

	"github.com/diegomagdaleno/tagger/converters"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func findFileProperties(slice []converters.FileProperties, val converters.FileProperties) (int, bool) {
	for i, item := range slice {
		if reflect.DeepEqual(item, val) {
			return i, true
		}
	}
	return -1, false
}

func SearchForFilesWithTags(files []converters.FileProperties, tag string) []converters.FileProperties {
	var filesEditable []converters.FileProperties
	for i := range files {
		tags := files[i].Tags
		for range tags {
			_, found := findFileProperties(filesEditable, files[i])
			if stringInSlice(tag, tags) && !found {
				filesEditable = append(filesEditable, files[i])
			}
		}
	}
	return filesEditable
}

func SearchForFilesWithTagsExclusively(files []converters.FileProperties, tag string) []converters.FileProperties {
	var filesEditable []converters.FileProperties
	for i := range files {
		tags := files[i].Tags
		for range tags {
			_, found := findFileProperties(filesEditable, files[i])
			if stringInSlice(tag, tags) && !found && !(len(files[i].Tags) > 1) {
				filesEditable = append(filesEditable, files[i])
			}
		}
	}
	return filesEditable
}
