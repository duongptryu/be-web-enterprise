package ginexport

import (
	"archive/zip"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"os"
	"web/common"
	component "web/components"
)

const Dir = "./assets/"

func ExportDocs(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		files, err := ioutil.ReadDir(Dir)
		if err != nil {
			panic(common.ErrInternal(err))
		}

		ar := zip.NewWriter(c.Writer)

		for _, file := range files {
			file1, _ := os.Open(Dir + file.Name())
			f1, _ := ar.Create(file.Name())
			io.Copy(f1, file1)
		}

		c.Writer.Header().Set("Content-type", "application/octet-stream")
		c.Writer.Header().Set("Content-Disposition", "attachment; filename=docs.zip")
		ar.Close()
	}
}
