package studentRepository

import (
	"gymn/internal/database"
	"gymn/internal/models"
	"gymn/internal/utils"
)

func GetAllStudents() ([]*models.Student, *utils.Error) {
	var students []*models.Student

	query := database.DB.Order("created_at desc")

	if err := query.Find(&students).Error; err != nil {
		return nil, utils.Throw(err.Error(), 400)
	}

	return students, nil
}
