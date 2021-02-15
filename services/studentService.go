package services

import (
	"errors"
	"fmt"
	"placement/models"

	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	Register(student *models.Student) (*models.Student, error)
	FindOneStudent(query bson.M) (*models.Student, error)
	LoginStudent(email, password string) (*models.Student, error)
	FindStudentByID(id string) (*models.Student, error)
	GetAllStudents(limit, skip *int64) []*models.Student
}

type studentService struct {
}

// NewStudentService method to create a new StudentService
func NewStudentService() StudentService {
	return &studentService{}
}

// Register method to register a new student
func (s *studentService) Register(student *models.Student) (*models.Student, error) {
	alreadyExistingStudent, _ := s.FindOneStudent(bson.M{operator.Or: []bson.M{{"email": student.Email}, {"UINNumber": student.UINNumber}}})
	if alreadyExistingStudent != nil {
		return nil, errors.New("email or uin already in use")
	}
	hashPassword, err := hash.HashPassword(student.Password)
	if err != nil {
		return nil, err
	}
	student.Password = hashPassword
	err = mgm.Coll(student).Create(student)
	if err != nil {
		return nil, err
	}
	return student, nil
}

// FindOneStudent find one student by a query
func (s *studentService) FindOneStudent(query bson.M) (*models.Student, error) {
	alreadyExistingStudent := &models.Student{}
	err := mgm.Coll(alreadyExistingStudent).First(query, alreadyExistingStudent)
	if err != nil {
		return nil, err
	}
	return alreadyExistingStudent, nil
}

func (s *studentService) LoginStudent(email, password string) (*models.Student, error) {
	student, _ := s.FindOneStudent(bson.M{"email": email})
	if student == nil {
		return nil, fmt.Errorf("UnAuthorize")
	}
	if hash.CheckPasswordHash(password, student.Password) {
		return nil, fmt.Errorf("UnAuthorize")
	}
	return student, nil
}

func (s *studentService) FindStudentByID(id string) (*models.Student, error) {
	student := &models.Student{}
	pid, err := student.PrepareID(id)
	if err != nil {
		return nil, err
	}
	result := mgm.Coll(student).FindOne(mgm.Ctx(), bson.M{"_id": pid}, &options.FindOneOptions{
		Projection: bson.M{"password": false},
	})
	if result.Err() != nil {
		return nil, result.Err()
	}
	err = result.Decode(student)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (s *studentService) GetAllStudents(limit, skip *int64) []*models.Student {
	student := &models.Student{}
	result, err := mgm.Coll(student).Find(mgm.Ctx(), bson.M{}, &options.FindOptions{
		Projection: bson.M{"password": false},
		Limit:      limit,
		Skip:       skip,
	})
	if err != nil {
		return nil
	}
	students := make([]*models.Student, 10)
	if err = result.All(mgm.Ctx(), &students); err != nil {
		return nil
	}
	return students
}
