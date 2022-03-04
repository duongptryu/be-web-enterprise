package uploadbiz

import (
	"context"
	"fmt"
	"path/filepath"
	"time"
	"web/common"
	"web/modules/upload/uploadmodel"
	"web/modules/upload/uploadstore"
)

type uploadFileBiz struct {
	store uploadstore.FileStore
}

func NewUploadFileBiz(store uploadstore.FileStore) *uploadFileBiz {
	return &uploadFileBiz{
		store: store,
	}
}

func (biz *uploadFileBiz) UploadFileBiz(ctx context.Context, fileName string, fileSize int64) (*common.File, error) {
	fileExt := filepath.Ext(fileName)
	newFileName := fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	file := common.File{
		SQLModelCreate: common.SQLModelCreate{},
		Name:           fileName,
		Size:           int(fileSize),
		Ext:            filepath.Ext(fileName),
		Url:            uploadmodel.PathFile + newFileName,
	}

	if err := biz.store.CreateFile(ctx, &file); err != nil {
		//delete img on s3
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	return &file, nil
}
