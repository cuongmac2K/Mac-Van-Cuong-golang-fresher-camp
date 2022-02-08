package qlbhstorage

import (
	"awesomeProject/modules/qlbh/qlbhmodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *qlbhmodel.QlbhCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
