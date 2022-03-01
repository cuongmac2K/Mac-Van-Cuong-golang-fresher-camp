package rslikebiz

import (
	"context"
	"demo/common"
	restaurantlikemodel "demo/modules/restaurantlike/model"
)

type ListUserRestaurantStore interface {
	GetUserLikeRestaurant(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]common.SimpleUser, error)
}
type listUserLikeRestaurantBiz struct {
	store ListUserRestaurantStore
}

func NewListUserRestaurantBiz(store ListUserRestaurantStore) *listUserLikeRestaurantBiz {
	return &listUserLikeRestaurantBiz{store: store}
}
func (biz *listUserLikeRestaurantBiz) ListUser(ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	users, err := biz.store.GetUserLikeRestaurant(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)

	}
	return users, nil
}
