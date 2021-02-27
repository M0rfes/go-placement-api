package services

import (
	"placement/models"

	"github.com/kamva/mgm/v3"
)

type ApplicationService interface {
	CreateApplication(*models.Application) error
}

type applicationService struct{}

func NewApplicationService() ApplicationService {
	return &applicationService{}
}

func (a *applicationService) CreateApplication(application *models.Application) error {
	return mgm.Coll(application).Create(application)
}
