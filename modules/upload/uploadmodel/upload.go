package uploadmodel

import (
	"errors"
	"strings"
	"web/common"
)

const PathFile = "https://groupbar.me/assets/"

var extWhiteList = map[string]bool{".docx": true, ".pdf": true, ".jpg": true, ".jpeg": true, ".png": true, ".xlsx": true}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot save uploaded file",
		"ErrCannotSaveFile",
	)
}

func ValidateFileExt(ext string) error {
	if _, exist := extWhiteList[strings.ToLower(ext)]; !exist {
		return ErrExtFileInvalid
	}
	return nil
}

var (
	ErrFileTooLarge   = common.NewCustomError(errors.New("file too large"), "file too large", "ErrFileTooLarge")
	ErrExtFileInvalid = common.NewCustomError(errors.New("file extension invalid"), "file extension invalid", "ErrExtFileInvalid")
)
