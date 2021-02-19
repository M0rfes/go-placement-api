package services

import (
	"fmt"
	"placement/models"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var (
	companyHash HashService
)

func init() {
	companyHash = NewHashService(15)
}

// CompanyService intercase for company service
type CompanyService interface {
	RegisterCompany(company *models.Company) (*models.Company, error)
	LoginCompany(email, password string) (*models.Company, error)
	FindOneCompany(query *bson.M, opts ...*options.FindOneOptions) (*models.Company, error)
}

type companyService struct {
}

// NewCompanyService companyService constructor
func NewCompanyService() CompanyService {
	return &companyService{}
}

func (s *companyService) RegisterCompany(company *models.Company) (*models.Company, error) {
	oldCompany, _ := s.FindOneCompany(&bson.M{"email": company.Email})
	if oldCompany != nil {
		return nil, fmt.Errorf("email already in use")
	}
	hashPassword, err := companyHash.HashPassword(company.Password)
	if err != nil {
		return nil, err
	}
	company.Password = hashPassword
	err = mgm.Coll(company).Create(company)
	if err != nil {
		return nil, err
	}
	return oldCompany, nil
}

func (s *companyService) LoginCompany(email, password string) (*models.Company, error) {
	company, _ := s.FindOneCompany(&bson.M{"email": email})
	if company == nil {
		return nil, fmt.Errorf("UnAuthorize")
	}
	if companyHash.CheckPasswordHash(company.Password, password) {
		return nil, fmt.Errorf("UnAuthorize")
	}
	return company, nil
}

func (s *companyService) FindOneCompany(query *bson.M, opts ...*options.FindOneOptions) (*models.Company, error) {
	company := &models.Company{}
	result := mgm.Coll(company).FindOne(mgm.Ctx(), query, opts...)
	if result.Err() != nil {
		return nil, result.Err()
	}
	err := result.Decode(company)
	if err != nil {
		return nil, err
	}
	return company, nil
}

// func (s *companyService) UpdateCompany(company *models.Company) (*models.Company, error) {

// }
