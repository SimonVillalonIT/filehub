package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/SimonVillalonIT/filehub/pkg/config"
	"github.com/labstack/echo/v4"
)

type File struct {
	Name        string
	IsDirectory bool
	Size        int64
	CreatedAt   time.Time
}

type FilesHandler struct {
	Config *config.Config
}

func (f *FilesHandler) GetFiles(c echo.Context) error {
	entries, err := os.ReadDir(f.Config.DataFolder)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}
	var files []File

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			file := File{Name: entry.Name(), IsDirectory: entry.IsDir()}
			files = append(files, file)
			continue
		}
		file := File{Name: entry.Name(), IsDirectory: entry.IsDir(), Size: info.Size(), CreatedAt: info.ModTime()}
		files = append(files, file)
	}
	return c.JSON(http.StatusOK, map[string]any{"data": files})
}

func (f *FilesHandler) PostFiles(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create(filepath.Join(f.Config.DataFolder, file.Filename))
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}
	return c.JSON(http.StatusAccepted, map[string]any{"success": true})
}
