package ginupload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"web/common"
	component "web/components"
	"web/modules/upload/uploadbiz"
	"web/modules/upload/uploadmodel"
	"web/modules/upload/uploadstore"
)

func Upload(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("upload")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := uploadmodel.ValidateFileExt(filepath.Ext(fileHeader.Filename)); err != nil {
			panic(err)
		}

		fileStore := uploadstore.NewSQLStore(appCtx.GetDatabase())
		biz := uploadbiz.NewUploadFileBiz(fileStore)
		f, err := biz.UploadFileBiz(c.Request.Context(), fileHeader.Filename, fileHeader.Size)

		//store image into own system
		c.SaveUploadedFile(fileHeader, fmt.Sprintf("./assets/%s", f.Name))

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.NewSimpleSuccessResponse(f))
	}
}
