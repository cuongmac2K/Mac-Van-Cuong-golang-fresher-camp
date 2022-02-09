package qlbhbiz

import (
	"awesomeProject/modules/qlbh/qlbhmodel"
	"context"
	"errors"
)

type CreateqlbhStore interface {
	Create(ctx context.Context, data *qlbhmodel.QlbhCreate) error
}

type creaqlbhBiz struct {
	store CreateqlbhStore
}

func NewCreateqlbhbiz(store CreateqlbhStore) *creaqlbhBiz {
	return &creaqlbhBiz{store: store}
}
func (biz *creaqlbhBiz) CreateRestaurant(ctx context.Context, data *qlbhmodel.QlbhCreate) error {
	if data.Name == "" {
		return errors.New("ten k co")
	}
	err := biz.store.Create(ctx, data)
	return err
}
