package converters

//#cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "../bridge/bridge.m"
// #include "../bridge/ext/NSTaggerURL.m"
import "C"
import "unsafe"

func NSStringToCString(s *C.NSString) *C.char { return C.NSStringToCString(s) }

func NSStringToGoString(s *C.NSString) string { return C.GoString(NSStringToCString(s)) }

func GoInt(i *C.NSNumber) int { return int(C.NSNumberToInt(i)) }

func GetNSArrayLenght(arr *C.NSArray) uint { return uint(C.NSArrayLen(arr)) }

func GetNSArrayItem(arr *C.NSArray, i uint) unsafe.Pointer {
	return C.NSArrayItem(arr, C.ulong(i))
}

func fromTGFilePropertiesToFileProperties(filePropertiesPtr *C.FileProperties) fileProperties {

	var tags []string

	filePropertiesFromC := *C.filePropertiesData(filePropertiesPtr)

	name := NSStringToGoString(filePropertiesFromC.name)

	lenght := GetNSArrayLenght(filePropertiesFromC.tags)

	for i := uint(0); i < lenght; i++ {
		tag := (*C.NSString)(GetNSArrayItem(filePropertiesFromC.tags, i))
		tagString := NSStringToGoString(tag)
		tags = append(tags, tagString)
	}

	filePropertiesGo := fileProperties{
		Name: name,
		Tags: tags,
	}

	return filePropertiesGo
}

func GoArrayWithFilePropertyObjects(arr *C.NSArray) []fileProperties {
	var files []fileProperties
	lenght := GetNSArrayLenght(arr)

	for i := uint(0); i < lenght; i++ {
		file := (*C.FileProperties)(GetNSArrayItem(arr, i))
		fileObject := fromTGFilePropertiesToFileProperties(file)
		files = append(files, fileObject)
	}
	return files
}

func GetFinalArrayOfFiles(path string) []fileProperties {
	return GoArrayWithFilePropertyObjects(C.getFilesWithCertainMacOSTag(C.CString(path)))
}
