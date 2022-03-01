package citiesbiz

import (
	"context"
	"demo/common"
	"demo/modules/cities/citiesmodel"
)

type getCitiesStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*citiesmodel.Cities, error)
}
type getCitiesbiz struct {
	store getCitiesStore
}

func NewgGetCities(store getCitiesStore) *getCitiesbiz {
	return &getCitiesbiz{store: store}
}
func (biz *getCitiesbiz) GetCities(ctx context.Context, id int) (*citiesmodel.Cities, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(citiesmodel.EntityName, err)
		}
	}

	if data.Id == 0 {
		return nil, common.ErrEntityDeleted(citiesmodel.EntityName, err)
	}
	return data, err
}
