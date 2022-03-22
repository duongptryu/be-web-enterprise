package uploadbiz

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"web/common"
	"web/modules/acayear/acayearstore"
	"web/modules/upload/uploadmodel"
	"web/modules/upload/uploadstore"
)

type uploadFileBiz struct {
	store        uploadstore.FileStore
	acaYearStore acayearstore.AcademicYearStore
}

func NewUploadFileBiz(store uploadstore.FileStore, acaYearStore acayearstore.AcademicYearStore) *uploadFileBiz {
	return &uploadFileBiz{
		store:        store,
		acaYearStore: acaYearStore,
	}
}

func (biz *uploadFileBiz) UploadFileBiz(ctx context.Context, fileName string, fileSize int64) (*common.File, error) {
	acaYear, err := biz.acaYearStore.FindAcaYear(ctx, map[string]interface{}{"status": true})
	if err != nil {
		return nil, err
	}

	err = common.CheckFolder(fmt.Sprintf("./assets/%v", acaYear.Id))
	if err != nil {
		panic(err)
	}

	fileExt := filepath.Ext(fileName)
	fileName = strings.ReplaceAll(fileName, " ", "_")
	fileName = strings.ReplaceAll(fileName, "/", "_")
	lastDotIndex := strings.LastIndex(fileName, ".")
	newFileName := fmt.Sprintf("%s-%d%s", fileName[:lastDotIndex], time.Now().Nanosecond(), fileExt)

	file := common.File{
		SQLModelCreate: common.SQLModelCreate{},
		Name:           newFileName,
		NameOrigin:     fileName,
		Size:           float64(fileSize),
		Ext:            filepath.Ext(fileName),
		Url:            fmt.Sprintf("%s%v/%s", uploadmodel.PathFile, acaYear.Id, newFileName),
		Folder:         strconv.Itoa(acaYear.Id),
	}

	if err := biz.store.CreateFile(ctx, &file); err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	return &file, nil
}

func (biz *uploadFileBiz) UploadFileCommonBiz(ctx context.Context, fileName string, fileSize int64) (*common.File, error) {
	err := common.CheckFolder("./assets/common")
	if err != nil {
		panic(err)
	}

	fileExt := filepath.Ext(fileName)
	fileName = strings.ReplaceAll(fileName, " ", "_")
	fileName = strings.ReplaceAll(fileName, "/", "_")
	lastDotIndex := strings.LastIndex(fileName, ".")
	newFileName := fmt.Sprintf("%s-%d%s", fileName[:lastDotIndex], time.Now().Nanosecond(), fileExt)

	file := common.File{
		SQLModelCreate: common.SQLModelCreate{},
		Name:           newFileName,
		NameOrigin:     fileName,
		Size:           float64(fileSize),
		Ext:            filepath.Ext(fileName),
		Url:            uploadmodel.PathFile + "common/" + newFileName,
		Folder:         "common",
	}

	return &file, nil
}
