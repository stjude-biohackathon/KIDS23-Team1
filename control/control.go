package control

import (
	"fileFlipper/model"
	"strings"

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
func Submit_order(user_id int, sample_id int, analysis_id int, order_name string, db *gorm.DB) string {
	// create order
	order := model.Order{
		Order_name:    order_name,
		UserRefer:     uint(user_id),
		SampleRefer:   uint(sample_id),
		AnslysisRefer: uint(analysis_id),
	}

	print(order.Order_name)

	// find all sample files on sample id
	var sample_files []model.SampleFile

	if err := db.Where("sample_refer = ?", sample_id).Find(&sample_files).Error; err != nil {
		return "error finding sample files"
	}

	pending_files := []model.SampleFile{}
	pending_files_name := []string{}
	msg := "Order scheduled for processing"

	// loop through sample files and check status
	for _, sample_file := range sample_files {
		if sample_file.File_state != "Ready" && sample_file.File_state != "Rehydrating" {
			pending_files = append(pending_files, sample_file)
			pending_files_name = append(pending_files_name, sample_file.File_name)
		}
	}

	// check if there are any pending files
	if len(pending_files) > 0 {
		// do something to initiate file rehydration
		// loop through pending files and check status
		for _, pending_file := range pending_files {
			print(pending_file.ID)
			// update file state to rehydrating by give file id
			db.Model(&model.SampleFile{}).Where("id = ?", pending_file.ID).Update("file_state", "Rehydrating")
		}
		msg = strings.Join(pending_files_name, " ")
	}

	return msg
}

func Cancel_order(user_id int, order_id int, db *gorm.DB) error {
	if err := db.Model(&model.Order{}).Where("id = ?", order_id).Update("order_state", "cancelled").Error; err != nil {
		return err
	}
	return nil
}

// to be implemented, Out of scope for now
func Modify_order(order_id int) {
	print("Hello World fronm order controller!")
}

func MoveFileFromColdToHot() {
	print("Hello World fronm order controller!")
}

// return all analyses
func Get_analyses(db *gorm.DB) ([]model.Analysis, error) {
	// get all analyses from db
	var analyses []model.Analysis
	if err := db.Find(&analyses).Error; err != nil {
		return nil, err
	}
	return analyses, nil
}
