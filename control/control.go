package control

import (
	"fileFlipper/model"

	"gorm.io/gorm"
)

func Get_user_by_id(user_id string, db *gorm.DB) (*model.User, error) {
	var user model.User
	if err := db.First(&user, user_id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// get sample related info
func Get_sample_and_related_files(sample_id string, db *gorm.DB) (*model.Sample, error) {
	// return sample with related files
	var sample model.Sample
	if err := db.Preload("SampleFiles").First(&sample, sample_id).Error; err != nil {
		return nil, err
	}
	return &sample, nil
}

// get order related info
func Get_orders_by_user_id(user_id int, db *gorm.DB) ([]model.Order, error) {
	var orders []model.Order
	if err := db.Where("user_refer = ?", user_id).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// order operations
func Submit_order(user_id int, sample_id int, order_id int) {
	print("Hello World fronm order controller!")
}

func Cancel_order(user_id int, order_id int, db *gorm.DB) error {
	if err := db.Model(&model.Order{}).Where("id = ?", order_id).Update("order_state", "cancelled").Error; err != nil {
		return err
	}
	return nil
}

func Modify_order(order_id int) {
	print("Hello World fronm order controller!")
}
