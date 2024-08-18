package models

import (
	"time"

	"github.com/pskclub/mine-core/utils"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string         `json:"id" gorm:"column:id;primary_key"`
	CreatedAt *time.Time     `json:"created_at" gorm:"column:CREATED_AT"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"column:UPDATED_AT"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:DELETED_AT"`
}

func NewBaseModel() BaseModel {
	return BaseModel{
		ID:        utils.GetUUID(),
		CreatedAt: utils.GetCurrentDateTime(),
		UpdatedAt: utils.GetCurrentDateTime(),
	}
}

type BaseModelHardDelete struct {
	ID        string     `json:"id" gorm:"column:id;primary_key"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:CREATED_AT"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:UPDATED_AT"`

	CreatedByID string `json:"created_by_id" gorm:"column:CREATED_BY_ID"`
	CreatedBy   string `json:"created_by" gorm:"-"`
	UpdatedByID string `json:"updated_by_id" gorm:"column:UPDATED_BY_ID"`
	UpdatedBy   string `json:"updated_by" gorm:"-"`
}

func NewBaseModelHardDelete() BaseModelHardDelete {
	return BaseModelHardDelete{
		ID:        utils.GetUUID(),
		CreatedAt: utils.GetCurrentDateTime(),
		UpdatedAt: utils.GetCurrentDateTime(),
	}
}

type BaseModelNotDelete struct {
	ID        string     `json:"id" gorm:"column:id;primary_key"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:CREATED_AT"`

	CreatedByID string `json:"created_by_id" gorm:"column:CREATED_BY_ID"`
	CreatedBy   string `json:"created_by" gorm:"-"`
}

func NewBaseModelNotDelete() BaseModelNotDelete {
	return BaseModelNotDelete{
		ID:        utils.GetUUID(),
		CreatedAt: utils.GetCurrentDateTime(),
	}
}

type BaseModelStatic struct {
	ID        string         `json:"id" gorm:"column:id;primary_key"`
	CreatedAt *time.Time     `json:"created_at" gorm:"column:CREATED_AT"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:DELETED_AT"`

	CreatedByID string `json:"created_by_id" gorm:"column:CREATED_BY_ID"`
	CreatedBy   string `json:"created_by" gorm:"-"`
	DeletedByID string `json:"deleted_by_id" gorm:"column:DELETED_BY_ID"`
	DeletedBy   string `json:"deleted_by" gorm:"-"`
}

func NewBaseModelStatic() BaseModelStatic {
	return BaseModelStatic{
		ID:        utils.GetUUID(),
		CreatedAt: utils.GetCurrentDateTime(),
	}
}
