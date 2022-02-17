package citiesbiz

import (
	"context"
	"demo/common"
	"demo/modules/cities/citiesmodel"
)

type DeleteCitiesStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*citiesmodel.Cities, error)
	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
}

type deleteCitiesBiz struct {
	store DeleteCitiesStore
}

func NewDeleteCitiestBiz(store DeleteCitiesStore) *deleteCitiesBiz {
	return &deleteCitiesBiz{store: store}
}

func (biz *deleteCitiesBiz) DeleteCitiest(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(citiesmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(citiesmodel.EntityName, nil)
	}

	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(citiesmodel.EntityName, err)
	}

	return nil
}
