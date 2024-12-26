package initializers

import "github.com/tao73bot/A_simple_CRM/models"

func Migrate() {
	// Drop all tables first to avoid circular dependencies
	// DB.Migrator().DropTable(&models.Admin{})
	
	// Create tables in order of dependencies
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Lead{})
	DB.AutoMigrate(&models.Customer{})
	DB.AutoMigrate(&models.Interaction{})
}