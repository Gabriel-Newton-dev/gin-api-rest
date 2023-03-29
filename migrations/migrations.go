package migrations

import (
	"github.com/Gabriel-Newton-dev/gin-api-rest/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{}) //created table user
}
