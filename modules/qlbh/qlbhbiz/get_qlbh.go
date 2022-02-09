package qlbhbiz

import (
	"awesomeProject/modules/common"
	"awesomeProject/modules/qlbh/qlbhmodel"
	"context"
)

type GetRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*qlbhmodel.Nhanvien, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(ctx context.Context, id int) (*qlbhmodel.Nhanvien, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(qlbhmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(qlbhmodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(qlbhmodel.EntityName, nil)
	}

	return data, err
}
