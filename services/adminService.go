package services

import (
	"fmt"
	"placement/models"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	adminHashService HashService
	studentS         StudentService
)

func init() {
	adminHashService = NewHashService(15)
	studentS = NewStudentService()
}

type AdminService interface {
	LoginAdmin(password string) (*models.Admin, error)
	ToggleAproven(studentID string) error
}

type adminService struct{}

func NewAdminService() AdminService {

	return &adminService{}
}

func (s *adminService) LoginAdmin(password string) (*models.Admin, error) {
	admin := &models.Admin{}
	_ = mgm.Coll(admin).First(bson.M{"username": "admin"}, admin)
	if !adminHashService.CheckPasswordHash(admin.Password, password) {
		return nil, fmt.Errorf("UnAuthorize")
	}
	return admin, nil
}

func (s *adminService) ToggleAproven(studentID string) error {
	student, err := studentS.FindStudentByID(studentID)
	if err != nil {
		return err
	}
	student.Approved = !student.Approved
	err = studentS.UpdateLoggedInStudent(student)
	if err != nil {
		return err
	}
	return nil
}
