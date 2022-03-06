package uploadmodel

import (
	"errors"
	"web/common"
)

const PathFile = "https://groupbar.me/assets/"

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot save uploaded file",
		"ErrCannotSaveFile",
	)
}

var (
	ErrFileTooLarge = common.NewCustomError(errors.New("file too large"), "file too large", "ErrFileTooLarge")
)
