package citiesstorage

import (
	"context"
	"demo/common"
	"demo/modules/cities/citiesmodel"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *citiesmodel.CitiesUpdate,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
