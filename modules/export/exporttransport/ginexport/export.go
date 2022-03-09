package ginexport

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"web/common"
	component "web/components"
	"web/modules/acayear/acayearstore"
	"web/modules/export/exportbiz"
	"web/modules/export/exportmodel"
	"web/modules/idea/ideastore"
)

func ExportIdeaToCsv(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		acaYearId, err := strconv.Atoi(c.Query("aca_year_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data exportmodel.Export
		data.AcaYearId = acaYearId

		store := ideastore.NewSQLStore(appCtx.GetDatabase())
		acaYearStore := acayearstore.NewSQLStore(appCtx.GetDatabase())
		biz := exportbiz.NewExportIdeaBiz(store, acaYearStore)

		result, err := biz.ExportIdeaBiz(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		downloadName := fmt.Sprintf("ideas-%s-%v.xlsx", data.NameAcaYear, time.Now().Unix())
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Disposition", "attachment; filename="+downloadName)
		c.Data(http.StatusOK, "application/octet-stream", result.Bytes())
	}
}
