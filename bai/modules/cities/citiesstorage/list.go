package citiesstorage

import (
	"demo/common"
	"demo/modules/cities/citiesmodel"
)

func (s *sqlStore) ListDataByCondition(
	condition map[string]interface{},
	filter *citiesmodel.Filter,
	paging *common.Paging,
	morekeys ...string,
) ([]citiesmodel.Cities, error) {
	var result []citiesmodel.Cities

	db := s.db
	for i := range morekeys {
		db = db.Preload(morekeys[i])
	}
	db = db.Table(citiesmodel.Cities{}.TableName()).Where(condition).Where("status in (1)")
	if v := filter; v != nil {
		if v.CitiesId > 0 {
			db = db.Where("id=?", v.CitiesId)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if err := db.
		Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
