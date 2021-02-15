package services

import (
	"errors"
	"placement/models"

	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"gopkg.in/mgo.v2/bson"
)

var (
	hash HashService
)

func init() {
	hash = &hashService{14}
}

// StudentService interface
type StudentService interface {
	Register(student *models.Student) error
	FindOneStudent(query bson.M, student *models.Student) *models.Student
}

type studentService struct {
}

// NewStudentService method to create a new StudentService
func NewStudentService() StudentService {
	return &studentService{}
}

// Register method to register a new student
func (s *studentService) Register(student *models.Student) error {
	alreadyExistingStudent := s.FindOneStudent(bson.M{operator.Or: []bson.M{{"email": student.Email}, {"UINNumber": student.UINNumber}}}, student)
	if alreadyExistingStudent.Email != "" {
		return errors.New("student already in existence")
	}
	hashPassword, err := hash.HashPassword(student.Password)
	if err != nil {
		return err
	}
	student.Password = hashPassword
	return mgm.Coll(student).Create(student)
}

// FindOneStudent find one student by a query
func (s *studentService) FindOneStudent(query bson.M, student *models.Student) *models.Student {
	alreadyExistingStudent := &models.Student{}
	_ = mgm.Coll(student).First(query, alreadyExistingStudent)
	return alreadyExistingStudent
}
