package acayearbiz

import (
	"context"
	"time"
	"web/common"
	"web/modules/acayear/acayearmodel"
	"web/modules/acayear/acayearstore"
)

type findAcaYear struct {
	store acayearstore.AcademicYearStore
}

func NewFindAcaYear(store acayearstore.AcademicYearStore) *findAcaYear {
	return &findAcaYear{
		store: store,
	}
}

func (biz *findAcaYear) FindAcaYear(ctx context.Context) (*acayearmodel.AcademicYear, error) {
	result, err := biz.store.FindAcaYear(ctx, map[string]interface{}{"status": true})
	if err != nil {
		return nil, common.ErrCannotListEntity(acayearmodel.EntityName, err)
	}
	if result.Id == 0 {
		return nil, common.ErrDataNotFound(acayearmodel.EntityName)
	}

	return result, nil
}

func (biz *findAcaYear) CheckCanCommentOrPostIdea(ctx context.Context) (*acayearmodel.CheckStatusPostIdeaComment, error) {
	acaYear, err := biz.store.FindAcaYear(ctx, map[string]interface{}{"status": true})
	if err != nil {
		return nil, common.ErrCannotListEntity(acayearmodel.EntityName, err)
	}
	if acaYear.Id == 0 {
		return nil, common.ErrDataNotFound(acayearmodel.EntityName)
	}

	timeNow := time.Now()

	var result acayearmodel.CheckStatusPostIdeaComment
	if timeNow.Before(acaYear.FirstClosureDate) {
		result.CanPostIdea = true
	}
	if timeNow.Before(acaYear.FinalClosureDate) {
		result.CanPostComment = true
	}

	return &result, nil
}
