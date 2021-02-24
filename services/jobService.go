package services

import (
	"placement/models"

	"github.com/kamva/mgm/v3"
	o "github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
)

// JobService service for managing jobs.
type JobService interface {
	CreateJob(job *models.Job) (*models.Job, error)
	GetAllJobs() *[]*models.Job
	GetJobById(id string) (*models.Job, error)
	UpdateJob(job *models.Job) error
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

func (j *jobService) GetAllJobs() *[]*models.Job {
	jobs := &[]*models.Job{}
	companyCollectionName := mgm.Coll(&models.Company{}).Name()
	ctx := mgm.Ctx()
	result, err := mgm.Coll(&models.Job{}).Aggregate(ctx, bson.A{
		bson.M{o.Lookup: bson.M{
			"from":         companyCollectionName,
			"localField":   "companyId",
			"foreignField": "_id",
			"as":           "company",
		},
		},
		bson.M{
			"$unwind": "$company",
		},
		bson.M{
			"$project": bson.M{
				"companyId": false,
				"company": bson.M{
					"password": false,
				},
			},
		},
	})
	if err != nil {
		return jobs
	}
	result.All(ctx, jobs)
	return jobs
}

func (j *jobService) GetJobById(id string) (*models.Job, error) {
	job := &models.Job{}
	jobs := []*models.Job{}
	companyCollectionName := mgm.Coll(&models.Company{}).Name()
	ctx := mgm.Ctx()
	pid, err := job.PrepareID(id)
	if err != nil {
		return nil, err
	}
	result, err := mgm.Coll(job).Aggregate(ctx, bson.A{
		bson.M{"$match": bson.M{"_id": pid}},
		bson.M{o.Lookup: bson.M{
			"from":         companyCollectionName,
			"localField":   "companyId",
			"foreignField": "_id",
			"as":           "company",
		},
		},
		bson.M{
			"$unwind": "$company",
		},
		bson.M{
			"$project": bson.M{
				"companyId": false,
				"company": bson.M{
					"password": false,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	result.All(ctx, &jobs)
	return jobs[0], nil
}

func (j *jobService) UpdateJob(job *models.Job) error {
	return mgm.Coll(job).Update(job)
}
