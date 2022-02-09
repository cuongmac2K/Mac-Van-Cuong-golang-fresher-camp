package compoment

import "gorm.io/gorm"

type AppContext interface {
	GetmainDBConnection() *gorm.DB
}
type appCtx struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{db: db}
}
func (ctx *appCtx) GetmainDBConnection() *gorm.DB {

	return ctx.db
}
