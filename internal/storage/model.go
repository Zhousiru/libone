package storage

import (
	"github.com/spf13/viper"
)

const (
	TypeFolder = "folder"
	TypeFile   = "file"
)

type Folder struct {
	path              string
	absPath           string
	base              string
	absParentFolder   string
	absParentPropPath string
	Prop              *viper.Viper
}

type ListItem struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
	FileType string `json:"file_type"`
}
