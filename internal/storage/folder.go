package storage

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Zhousiru/libone/internal/util"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// NewFolder create a Folder structure by path.
// If the path doesn't exist, NewFolder will create it.
func NewFolder(path string) (*Folder, error) {
	f := new(Folder)

	f.absPath = util.GetAbsPath(path)

	absDataPath, err := filepath.Abs(viper.GetString("data-path"))
	if err != nil {
		return nil, err
	}
	if !strings.HasPrefix(f.absPath, absDataPath) {
		return nil, errors.New("the specified path is forbidden")
	}

	f.path, err = filepath.Rel(absDataPath, f.absPath)
	if err != nil {
		return nil, err
	}
	if !util.PathExist(f.absPath) {
		os.MkdirAll(f.absPath, os.FileMode(0666))
	}

	f.base = filepath.Base(f.absPath)

	if f.absPath != absDataPath {
		f.absParentFolder = filepath.Dir(f.absPath)
		f.absParentPropPath = filepath.Join(f.absParentFolder, ".lo")

		if !util.PathExist(f.absParentPropPath) {
			os.MkdirAll(f.absParentPropPath, os.FileMode(0666))
		}

		prop := viper.New()
		prop.SetConfigType("yaml")
		prop.AddConfigPath(f.absParentPropPath)
		prop.SetConfigName(f.base)
		prop.ReadInConfig()

		absSelfPropPath := filepath.Join(f.absParentPropPath, f.base+".yaml")
		if !util.PathExist(absSelfPropPath) {
			f, err := os.Create(absSelfPropPath)
			if err != nil {
				return nil, err
			}
			defer f.Close()
		}

		prop.WatchConfig()
		prop.OnConfigChange(func(e fsnotify.Event) {
			f.Prop.ReadInConfig()
		})
		f.Prop = prop
	}

	return f, nil
}

// List returns a slice of files.
func (f *Folder) List() ([]ListItem, error) {
	var list []ListItem

	fis, err := ioutil.ReadDir(f.absPath)
	if err != nil {
		return nil, err
	}

	for _, fi := range fis {
		item := ListItem{}
		item.Filename = fi.Name()
		item.Path = filepath.Join(f.path, fi.Name())
		if fi.IsDir() {
			item.FileType = TypeFolder
		} else {
			item.FileType = TypeFile
		}

		list = append(list, item)
	}

	return list, nil
}
