package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetCurrentPath() string {
	exePath, err := os.Executable()

	if err != nil {
		fmt.Println(err)
		return ""
	}

	path, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return path
}
