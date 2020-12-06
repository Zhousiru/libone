package api

import (
	"io/ioutil"

	"github.com/Zhousiru/libone/internal/storage"

	"github.com/Zhousiru/libone/internal/util"

	"github.com/gin-gonic/gin"
)

// GetFile returns a file.
func GetFile(c *gin.Context) {
	absPath := util.GetAbsPath(c.Param("path"))
	if !util.PathExist(absPath) {
		resp(c, respStatusErr, "the specified path does not exist", nil)
		return
	}
	if util.IsFolder(absPath) {
		resp(c, respStatusErr, "the specified path is invalid", nil)
		return
	}
	c.File(absPath)
}

// ListFile returns a list of files in the specified folder.
func ListFile(c *gin.Context) {
	path := c.Query("path")
	absPath := util.GetAbsPath(path)
	if !util.PathExist(absPath) {
		resp(c, respStatusErr, "the specified path does not exist", nil)
		return
	}
	if !util.IsFolder(absPath) {
		resp(c, respStatusErr, "the specified path is invalid", nil)
		return
	}

	folder, err := storage.NewFolder(path)
	if err != nil {
		resp(c, respStatusErr, err.Error(), nil)
		return
	}

	li, err := folder.List()
	if err != nil {
		resp(c, respStatusErr, err.Error(), nil)
		return
	}

	resp(c, respStatusOK, "", li)
}

// UpdateFile updates a file to the specified folder.
func UpdateFile(c *gin.Context) {
	err := c.Request.ParseMultipartForm(50 << 20)
	if err != nil {
		panic(err)
	}

	files := c.Request.MultipartForm.File["file"]
	path := c.Query("path")
	absPath := util.GetAbsPath(path)
	if !util.PathExist(absPath) {
		resp(c, respStatusErr, "the specified path does not exist", nil)
		return
	}
	if !util.IsFolder(absPath) {
		resp(c, respStatusErr, "the specified path is invalid", nil)
		return
	}

	folder, err := storage.NewFolder(path)
	if err != nil {
		resp(c, respStatusErr, err.Error(), nil)
		return
	}

	for _, f := range files {
		file, err := f.Open()
		if err != nil {
			resp(c, respStatusErr, err.Error(), nil)
			return
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			resp(c, respStatusErr, err.Error(), nil)
			return
		}

		folder.WriteFile(f.Filename, data)
	}

	resp(c, respStatusOK, "", nil)
}
