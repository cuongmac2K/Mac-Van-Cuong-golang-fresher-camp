package qlbhstorage

import (
	"awesomeProject/modules/common"
	"awesomeProject/modules/qlbh/qlbhmodel"
	"context"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *qlbhmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]qlbhmodel.Nhanvien, error) {
	var result []qlbhmodel.Nhanvien

	db := s.db
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	db = db.Table(qlbhmodel.Nhanvien{}.TableName()).Where(conditions)
	if v := filter; v != nil {
		if v.DateOfBirth != "" {
			db = db.Where("DateOfBirth = ?", v.DateOfBirth)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
