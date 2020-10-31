package lib

import (
	"fmt"
	"os"
)

func GetCurrentDirectory() string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return currentDir
}
