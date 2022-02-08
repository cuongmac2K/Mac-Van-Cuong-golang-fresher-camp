package qlbhbiz

import (
	"awesomeProject/modules/common"
	"awesomeProject/modules/qlbh/qlbhmodel"
	"context"
)

type ListQlBHStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *qlbhmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]qlbhmodel.Nhanvien, error)
}

type listQLBHBiz struct {
	store ListQlBHStore
}

func NewlistQLBHBiz(store ListQlBHStore) *listQLBHBiz {
	return &listQLBHBiz{store: store}
}
func (biz *listQLBHBiz) ListQLBH(
	ctx context.Context,
	filter *qlbhmodel.Filter,
	paging *common.Paging) ([]qlbhmodel.Nhanvien, error) {

	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)
	return result, err
}
