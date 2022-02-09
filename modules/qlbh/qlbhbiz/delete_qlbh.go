package qlbhbiz

import (
	"awesomeProject/modules/common"
	"awesomeProject/modules/qlbh/qlbhmodel"
	"context"
)


type DeleteQLBHStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*qlbhmodel.Nhanvien.Restaurant, error)
	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
}

type deleteQLBHBiz struct {
	store DeleteQLBHStore
}


func NewDeleteRestaurantBiz(store DeleteQLBHStore) *deleteQLBHBiz {
	return &deleteQLBHBiz{store: store}
}

func (biz *deleteQLBHBiz) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(qlbhmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(qlbhmodel.EntityName, nil)
	}

	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(qlbhmodel.EntityName, err)
	}

	return nil
}
