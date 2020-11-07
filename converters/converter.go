package converters

//#cgo CFLAGS: -x objective-c -I /usr/local/include
// #cgo LDFLAGS: -framework Foundation -L/usr/local/lib -ltags
// #include "../bridge/bridge.m"
// #include <libtags/NSURLTagger.h>
// #include <libtags/GetHelpers.h>
// #include <libtags/TagComponents.h>
import "C"
import (
	"errors"
	"unsafe"
)

// NSStringToCString handles the conversion from an NSString object to a CString variable.
func NSStringToCString(s *C.NSString) *C.char { return C.NSStringToCString(s) }

// NSStringToGoString handles the conversion from NSString to our own GoString session
func NSStringToGoString(s *C.NSString) string { return C.GoString(NSStringToCString(s)) }

func GoStringToNSString(s string) *C.NSString { return C.cStringToNSString(C.CString(s)) }

// GoInt handles the conversion from NSNumber to Go's int
func GoInt(i *C.NSNumber) int { return int(C.NSNumberToInt(i)) }

// GetNSArrayLenght returns the number of elements stored inside an array as an uint
func GetNSArrayLenght(arr *C.NSArray) uint { return uint(C.NSArrayLen(arr)) }

// GetNSArrayItem returns certain value at an index
func GetNSArrayItem(arr *C.NSArray, i uint) unsafe.Pointer {
	return C.NSArrayItem(arr, C.ulong(i))
}

func fromTGTagComponentsToTagComponets(tagComponentsPtr *C.TagComponent) TagComponents {

	tagComponentsFromC := *C.tagComponentsData(tagComponentsPtr)

	name := NSStringToGoString(tagComponentsFromC.name)
	color := NSStringToGoString(tagComponentsFromC.color)

	TagComponentGo := TagComponents{
		Name:  name,
		Color: color,
	}

	return TagComponentGo
}

// fromTGFilePropertiesToFileProperties
// Handles the conversion from the Obj-C struct TGFileProperties to our own fileProperties
// type.
// How it works:
// First we get a TGFileProperties from the function filePropertiesData in the C file (This function is documented in that file)
// We convert the name to a Go name
// The rest of the code until filePropertiesGo variable declaration, converts the array in each object to a GoArray
// finally we init our own version of this object and we return it
func fromTGFilePropertiesToFileProperties(filePropertiesPtr *C.FileProperties) FileProperties {

	var tags []TagComponents

	filePropertiesFromC := *C.filePropertiesData(filePropertiesPtr)

	name := NSStringToGoString(filePropertiesFromC.name)

	lenght := GetNSArrayLenght(filePropertiesFromC.tags)

	for i := uint(0); i < lenght; i++ {
		tag := (*C.TagComponent)(GetNSArrayItem(filePropertiesFromC.tags, i))
		tagString := fromTGTagComponentsToTagComponets(tag)
		tags = append(tags, tagString)
	}

	filePropertiesGo := FileProperties{
		Name: name,
		Tags: tags,
	}

	return filePropertiesGo
}

func NSErrorlocalizedDescriptionToGoError(localDesc *C.NSString) error {
	errorString := NSStringToGoString(localDesc)
	if errorString != "" {
		return errors.New(errorString)
	}
	return nil
}

// GoArrayWithFilePropertyObjects translates the returned array from Obj-C to a Go array containing our fileProperties objects
func GoArrayWithFilePropertyObjects(arr *C.NSArray) []FileProperties {
	var files []FileProperties
	lenght := GetNSArrayLenght(arr)

	for i := uint(0); i < lenght; i++ {
		file := (*C.FileProperties)(GetNSArrayItem(arr, i))
		fileObject := fromTGFilePropertiesToFileProperties(file)
		files = append(files, fileObject)
	}
	return files
}

func GetFinalArrayOfFiles(path string) []FileProperties {
	return GoArrayWithFilePropertyObjects(C.getTagsOfFile(GoStringToNSString(path)))
}

func RemoveTag(tag string, path string) error {

	success := NSErrorlocalizedDescriptionToGoError(C.removeTagsForFile(C.CString(tag), C.CString(path)))

	return success

}
