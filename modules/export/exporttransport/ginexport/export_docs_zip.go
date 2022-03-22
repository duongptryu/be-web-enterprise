package ginexport

import (
	"archive/zip"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/acayear/acayearstore"
	"web/modules/export/exportbiz"
)

const Dir = "./assets"

func ExportDocs(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		acaYearId, err := strconv.Atoi(c.Query("aca_year_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		acaYearStore := acayearstore.NewSQLStore(appCtx.GetDatabase())
		biz := exportbiz.NewExportDocsZip(acaYearStore)

		dir, err := biz.ExportDocsZip(c.Request.Context(), acaYearId)
		if err != nil {
			panic(err)
		}

		files, err := ioutil.ReadDir(dir)
		if err != nil {
			panic(common.ErrInternal(err))
		}

		ar := zip.NewWriter(c.Writer)

		for _, file := range files {
			file1, err := os.Open(dir + file.Name())
			if err != nil {
				panic(common.ErrInternal(err))
			}
			f1, err := ar.Create(file.Name())
			if err != nil {
				panic(common.ErrInternal(err))
			}
			_, err = io.Copy(f1, file1)
			if err != nil {
				panic(common.ErrInternal(err))
			}
		}

		c.Writer.Header().Set("Content-type", "application/octet-stream")
		c.Writer.Header().Set("Content-Disposition", "attachment; filename=docs.zip")
		ar.Close()
	}
}
