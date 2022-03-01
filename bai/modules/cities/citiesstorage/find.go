package citiesstorage

import (
	"context"
	"demo/common"
	"demo/modules/cities/citiesmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	condition map[string]interface{},
	morekey ...string,
) (*citiesmodel.Cities, error) {
	var result *citiesmodel.Cities

	db := s.db

	for i := range morekey {
		db = db.Preload(morekey[i])
	}

	if err := db.Where(condition).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
	}
	return result, nil
}
