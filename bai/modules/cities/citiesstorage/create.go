package citiesstorage

import (
	"context"
	"demo/common"
	"demo/modules/cities/citiesmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *citiesmodel.CitiesCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
