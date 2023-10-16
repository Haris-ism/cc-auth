package postgre

import (
	"cc-auth/models"

	"github.com/sirupsen/logrus"
)

func (db *postgreDB) Insert(req models.ItemList) error {
	if err := db.postgre.Create(&req).Error; err != nil {
		logrus.Errorf("Failed to Insert Postgre, Err:", err)
		return err
	}
	return nil
}

func (db *postgreDB) Query() ([]models.ItemList, error) {
	result := []models.ItemList{}
	if err := db.postgre.Where("price > ?", 1).Find(&result).Error; err != nil {
		logrus.Errorf("Failed to Insert Postgre, Err:", err)
		return result, err
	}
	return result, nil
}
