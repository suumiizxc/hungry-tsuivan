package marketplace

import "github.com/suumiizxc/car-marketplace/config"

type CarMark struct {
	ID               uint64 `json:"id" gorm:"primary_key"`
	CarManufactoryID uint64 `json:"car_manufactory_id"`
	Name             string `json:"name"`
}

func (cm *CarMark) Create() error {
	err := config.DB.Create(&cm).Error
	return err
}

func (cm *CarMark) FindByID() (CarMark, error) {
	var cma CarMark
	if err := config.DB.Find(&cma, cm.ID).Error; err != nil {
		return CarMark{}, err
	}
	return cma, nil
}

func (cm *CarMark) FindByCMID() ([]CarMark, error) {
	var cmas []CarMark
	if err := config.DB.Where("car_manufactory_id = ?", cm.CarManufactoryID).Find(&cmas).Error; err != nil {
		return []CarMark{}, err
	}
	return cmas, nil
}
