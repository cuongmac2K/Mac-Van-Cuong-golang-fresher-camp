package rsstorage

import (
	"context"
	"demo/common"
	"demo/modules/restaurantlike/model"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
)

func (s *sqlStore) GetRestaurantLikes(ctx context.Context,
	ids []int) (map[int]int, error) {
	result := make(map[int]int)

	type sqlData struct {
		RestaurantId int `gorm:"column:restaurant_id;"`
		LikeCount    int `gorm:"column:count;"`
	}

	var listLike []sqlData

	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).
		Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.RestaurantId] = item.LikeCount
	}

	return result, nil
	//return nil, errors.New("cannot get likes")
}
func (s *sqlStore) GetUserLikeRestaurant(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]common.SimpleUser, error) {
	var result []restaurantlikemodel.Like

	db := s.db

	//for i := range moreKeys {
	//	db = db.Preload(moreKeys[i])
	//
	//}
	db = db.Preload("User")
	db = db.Table(restaurantlikemodel.Like{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.RestaurantId > 0 {
			db = db.Where("restaurant_id = ?", v.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("created_at <?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(result))
	for i, item := range result {
		users[i] = *result[i].User
		if i == (len(result) - 1) {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format("2006-01-02T15:04:05.999999-07:00"))))
			paging.NextCursor = cursorStr
		}
	}
	return users, nil
}
