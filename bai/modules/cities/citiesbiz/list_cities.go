package citiesbiz

import (
	"context"
	"demo/common"
	"demo/modules/cities/citiesmodel"
)

type ListCitiesStore interface {
	ListCitiesStore(ctx context.Context,
		conditions map[string]interface{},
		filter *citiesmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]citiesmodel.Cities, error)
}
type ListCitiesBiz struct {
	store ListCitiesStore
}

func NewListCities(store ListCitiesStore) *ListCitiesBiz {
	return &ListCitiesBiz{store: store}
}
func (biz *ListCitiesBiz) LisCities(
	ctx context.Context,
	filter *citiesmodel.Filter,
	paging *common.Paging,
) ([]citiesmodel.Cities, error) {
	result, err := biz.store.ListCitiesStore(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(citiesmodel.EntityName, err)
	}
	return result, err
}
