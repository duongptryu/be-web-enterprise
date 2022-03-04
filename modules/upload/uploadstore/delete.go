package uploadstore

import (
	"context"
	"web/common"
)

func (s *sqlStore) DeleteFile(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(common.File{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
