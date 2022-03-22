package ginupload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"web/common"
	component "web/components"
	"web/modules/acayear/acayearstore"
	"web/modules/upload/uploadbiz"
	"web/modules/upload/uploadmodel"
	"web/modules/upload/uploadstore"
)

func UploadCommon(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("upload")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := uploadmodel.ValidateFileExt(filepath.Ext(fileHeader.Filename)); err != nil {
			panic(err)
		}

		fileStore := uploadstore.NewSQLStore(appCtx.GetDatabase())
		acaYearStore := acayearstore.NewSQLStore(appCtx.GetDatabase())
		biz := uploadbiz.NewUploadFileBiz(fileStore, acaYearStore)
		f, err := biz.UploadFileCommonBiz(c.Request.Context(), fileHeader.Filename, fileHeader.Size)

		//store image into own system
		err = c.SaveUploadedFile(fileHeader, fmt.Sprintf("./assets/%s/%s", f.Folder, f.Name))
		if err != nil {
			panic(common.ErrInternal(err))
		}

		c.JSON(200, common.NewSimpleSuccessResponse(f))
	}
}
