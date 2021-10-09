package seed

import (
	"log"

	"github.com/SilverNate/crud/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Email:    "koga@gmail.com",
		Password: "password",
		Address:  "Jl.Manggis No.29",
	},
	models.User{
		Email:    "testing@gmail.com",
		Password: "password",
		Address:  "Jl.Apel No.99",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

	}
}
