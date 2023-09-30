package postgre

import (
	auth "cc-auth/controllers/models"
	dbs "cc-auth/databases/postgresql/models"
	"cc-auth/models"
	"cc-auth/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	postgreDB struct {
		postgre *gorm.DB
	}
	PostgreInterface interface {
		Insert(req models.ItemList) error
		Query() ([]models.ItemList, error)
		Register(req auth.Credentials)error
	}
)

func InitPostgre() PostgreInterface {
	host := utils.GetEnv("POSTGRE")
	db, err := gorm.Open(postgres.Open(host), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Errorf("Failed to Init Postgre, Err:", err)
	} else {
		logrus.Printf("Init Postgre Success")
	}
	db.AutoMigrate(&models.ItemList{},&dbs.Credentials{})

	return &postgreDB{
		postgre: db,
	}
}
