package citiesstorage

import (
	"context"
	"demo/common"
	"demo/modules/cities/citiesmodel"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func (s *sqlStore) ListCitiesStore(ctx context.Context, conditions map[string]interface{}, filter *citiesmodel.Filter, paging *common.Paging, moreKeys ...string) ([]citiesmodel.Cities, error) {
	//TODO implement me
	panic("implement me")
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
