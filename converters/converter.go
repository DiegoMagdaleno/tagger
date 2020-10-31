package converters

//#cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "../bridge/bridge.m"
import "C"
import "unsafe"

func NSStringToCString(s *C.NSString) *C.char { return C.NSStringToCString(s) }

func NSStringToGoString(s *C.NSString) string { return C.GoString(NSStringToCString(s)) }

func GoInt(i *C.NSNumber) int { return int(C.NSNumberToInt(i)) }

func GetNSArrayLenght(arr *C.NSArray) uint { return uint(C.NSArrayLen(arr)) }

func GetNSArrayItem(arr *C.NSArray, i uint) unsafe.Pointer {
	return C.NSArrayItem(arr, C.ulong(i))
}

func GoArrayWithFiles(arr *C.NSArray) []string {
	var paths []string
	lenght := GetNSArrayLenght(arr)

	for i := uint(0); i < lenght; i++ {
		path := (*C.NSString)(GetNSArrayItem(arr, i))
		pathString := NSStringToGoString(path)
		paths = append(paths, pathString)
	}

	return paths
}

func GetFinalArrayOfFiles(path string) []string {
	return GoArrayWithFiles(C.getFilesWithCertainMacOSTag(C.CString(path)))
}
