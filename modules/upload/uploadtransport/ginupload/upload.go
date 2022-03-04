package ginupload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web/common"
	component "web/components"
	"web/modules/upload/uploadbiz"
	"web/modules/upload/uploadstore"
)

func Upload(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		//store image into own system
		c.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s", fileHeader.Filename))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		fileStore := uploadstore.NewSQLStore(appCtx.GetDatabase())
		biz := uploadbiz.NewUploadFileBiz(fileStore)
		img, err := biz.UploadFileBiz(c.Request.Context(), fileHeader.Filename, fileHeader.Size)

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.NewSimpleSuccessResponse(img))
	}
}
