package services

import (
	"placement/models"

	"github.com/kamva/mgm/v3"
)

// JobService service for managing jobs.
type JobService interface {
	CreateJob(job *models.Job) (*models.Job, error)
}

type jobService struct{}

func NewJobService() JobService {
	return &jobService{}
}

func (j *jobService) CreateJob(job *models.Job) (*models.Job, error) {
	err := mgm.Coll(job).Create(job)
	if err != nil {
		return nil, err
	}
	return job, nil
}
