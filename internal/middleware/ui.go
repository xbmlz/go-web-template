package middleware

import (
	"io/fs"
	"net/http"
	"path"

	"github.com/gin-contrib/static"
	"github.com/xbmlz/go-web-template/internal/logger"
	"github.com/xbmlz/go-web-template/ui"
)

type ServerFileSystemType struct {
	http.FileSystem
}

func (f ServerFileSystemType) Exists(prefix string, _path string) bool {
	file, err := f.Open(path.Join(prefix, _path))
	if file != nil {
		defer func(file http.File) {
			err = file.Close()
			if err != nil {
				logger.Error("file not found", err)
			}
		}(file)
	}
	return err == nil
}

func MustFS(dir string) (serverFileSystem static.ServeFileSystem) {
	sub, err := fs.Sub(ui.DistFS, path.Join("dist", dir))

	if err != nil {
		logger.Error(err)
		return
	}

	serverFileSystem = ServerFileSystemType{
		http.FS(sub),
	}
	return
}
