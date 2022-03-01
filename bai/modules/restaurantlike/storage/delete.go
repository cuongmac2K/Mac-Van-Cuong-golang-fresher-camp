package rsstorage

import (
	"context"
	"demo/common"
	restaurantlikemodel "demo/modules/restaurantlike/model"
)

func (s *sqlStore) Delete(ctx context.Context, userid, restaurantid int) error {
	db := s.db

	if err := db.Table(restaurantlikemodel.Like{}.TableName()).Where("user_id = ? and restaurant_id", userid, restaurantid).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
