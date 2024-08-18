package core

import (
	"context"

	"gorm.io/gorm"
)

type IContext interface {
	context.Context
	DB() *gorm.DB
}

type AppContext struct {
	context.Context
	db *gorm.DB
}

func NewContext(ctx context.Context, db *gorm.DB) IContext {
	return &AppContext{
		Context: ctx,
		db:      db,
	}
}

func (c *AppContext) DB() *gorm.DB {
	return c.db
}
