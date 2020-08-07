package util

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// PathExist check the existence of a path.
func PathExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}

// IsFolder checks if the path is a folder.
func IsFolder(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	return fi.IsDir()
}

// GetAbsPath returns the absolute path.
func GetAbsPath(path string) string {
	return filepath.Join(viper.GetString("data-path"), path)
}
