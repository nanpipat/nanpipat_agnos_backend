package migrations

import (
	"github.com/nanpipat/nanpipat-agnos-backend/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.PasswordLog{},
		// เพิ่มโมเดลอื่นๆ ที่ต้องการ migrate ที่นี่
	)
}
