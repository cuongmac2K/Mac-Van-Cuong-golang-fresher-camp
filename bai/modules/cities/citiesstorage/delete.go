package citiesstorage

import (
	"context"
	"demo/common"
	"demo/modules/cities/citiesmodel"
)

func (s *sqlStore) SoftDeleteData(
	ctx context.Context,
	id int,
) error {
	db := s.db

	if err := db.Table(citiesmodel.Cities{}.TableName()).
		Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
