package services

import (
	"errors"
	"fmt"
	"placement/models"

	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var (
	studentHash HashService
)

func init() {
	studentHash = NewHashService(14)
}

// StudentService interface
type StudentService interface {
	RegisterStudent(student *models.Student) (*models.Student, error)
	FindOneStudent(query *bson.M, opts ...*options.FindOneOptions) (*models.Student, error)
	LoginStudent(email, password string) (*models.Student, error)
	FindStudentByID(id string, opts ...*options.FindOneOptions) (*models.Student, error)
	GetAllStudents() *[]*models.Student
	UpdateLoggedInStudent(student *models.Student) error
	GetAllApprovedStudents() *[]*models.Student
	GetAllUnApprovedStudents() *[]*models.Student
	DeleteStudent(id string) error
}

type studentService struct {
}

// NewStudentService method to create a new StudentService
func NewStudentService() StudentService {
	return &studentService{}
}

// Register method to register a new student
func (s *studentService) RegisterStudent(student *models.Student) (*models.Student, error) {
	alreadyExistingStudent, _ := s.FindOneStudent(&bson.M{operator.Or: []bson.M{{"email": student.Email}, {"UINNumber": student.UINNumber}}})
	if alreadyExistingStudent != nil {
		return nil, errors.New("email or uin already in use")
	}
	hashPassword, err := studentHash.HashPassword(student.Password)
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
func (s *studentService) FindOneStudent(query *bson.M, opts ...*options.FindOneOptions) (*models.Student, error) {
	alreadyExistingStudent := &models.Student{}
	result := mgm.Coll(alreadyExistingStudent).FindOne(mgm.Ctx(), query, opts...)

	if err := result.Err(); err != nil {
		return nil, err
	}
	err := result.Decode(alreadyExistingStudent)
	if err != nil {
		return nil, err
	}
	return alreadyExistingStudent, nil
}

func (s *studentService) LoginStudent(email, password string) (*models.Student, error) {
	student, _ := s.FindOneStudent(&bson.M{"email": email})
	if student == nil {
		return nil, fmt.Errorf("UnAuthorize")
	}
	if !student.Approved {
		return nil, fmt.Errorf("Your account isn't approved yet")
	}
	if !studentHash.CheckPasswordHash(student.Password, password) {
		return nil, fmt.Errorf("UnAuthorize")
	}
	return student, nil
}

func (s *studentService) FindStudentByID(id string, opts ...*options.FindOneOptions) (*models.Student, error) {
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	student, err := s.FindOneStudent(&bson.M{"_id": pid}, opts...)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (s *studentService) GetAllStudents() *[]*models.Student {
	student := &models.Student{}
	result, err := mgm.Coll(student).Find(mgm.Ctx(), bson.M{}, &options.FindOptions{
		Projection: bson.M{"password": false},
	})
	students := &[]*models.Student{}
	if err != nil {
		return students
	}
	if err = result.All(mgm.Ctx(), students); err != nil {
		return students
	}
	return students
}

func (s *studentService) GetAllApprovedStudents() *[]*models.Student {
	student := &models.Student{}
	result, err := mgm.Coll(student).Find(mgm.Ctx(), bson.M{"approved": true}, &options.FindOptions{
		Projection: bson.M{"password": false},
	})
	students := &[]*models.Student{}
	if err != nil {
		return students
	}
	if err = result.All(mgm.Ctx(), students); err != nil {
		return students
	}
	return students
}

func (s *studentService) GetAllUnApprovedStudents() *[]*models.Student {
	student := &models.Student{}
	result, err := mgm.Coll(student).Find(mgm.Ctx(), bson.M{"approved": false}, &options.FindOptions{
		Projection: bson.M{"password": false},
	})
	students := &[]*models.Student{}
	if err != nil {
		return students
	}
	if err = result.All(mgm.Ctx(), students); err != nil {
		return students
	}
	return students
}

func (s *studentService) UpdateLoggedInStudent(student *models.Student) error {
	return mgm.Coll(student).Update(student)
}

func (s *studentService) DeleteStudent(id string) error {
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	student := &models.Student{}
	result := mgm.Coll(student).FindOneAndDelete(mgm.Ctx(), bson.M{"id": pid})
	if err := result.Err(); err != nil {
		return err
	}
	return nil
}
