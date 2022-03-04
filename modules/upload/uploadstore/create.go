package uploadstore

import (
	"context"
	"web/common"
)

func (s *sqlStore) CreateFile(ctx context.Context, data *common.File) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
