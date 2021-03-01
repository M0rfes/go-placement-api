package services

import (
	"fmt"
	"placement/models"

	"github.com/kamva/mgm/v3"
	o "github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	b "gopkg.in/mgo.v2/bson"
)

// JobService service for managing jobs.
type JobService interface {
	CreateJob(job *models.Job) (*models.Job, error)
	GetAllJobs() *[]*models.Job
	GetJobByID(id string) (*models.Job, error)
	UpdateJob(job *models.Job) error
	GetAllJobsForCompany(company string) *[]*models.Job
	FindOneJob(query *b.M, opts ...*options.FindOneOptions) (*models.Job, error)
	DeleteJob(id string) error
}

type jobService struct{}

// NewJobService constructor for JobService.
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

func (j *jobService) GetAllJobsForCompany(company string) *[]*models.Job {
	jobs := &[]*models.Job{}
	ctx := mgm.Ctx()
	pid, err := primitive.ObjectIDFromHex(company)
	if err != nil {
		return jobs
	}
	result, err := mgm.Coll(&models.Job{}).Find(ctx, bson.M{"companyId": pid}, &options.FindOptions{
		Projection: bson.M{"companyId": false, "company": false},
	})
	if err != nil {
		return jobs
	}
	result.All(ctx, jobs)
	return jobs
}

// GetJobByID to get a job by id.
func (j *jobService) GetJobByID(id string) (*models.Job, error) {
	jobs := []*models.Job{}
	companyCollectionName := mgm.Coll(&models.Company{}).Name()
	ctx := mgm.Ctx()
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result, err := mgm.Coll(&models.Job{}).Aggregate(ctx, bson.A{
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
	if len(jobs) == 0 {
		return nil, fmt.Errorf("no job found")
	}
	return jobs[0], nil
}

func (j *jobService) UpdateJob(job *models.Job) error {
	return mgm.Coll(job).Update(job)
}

func (j *jobService) FindOneJob(query *b.M, opts ...*options.FindOneOptions) (*models.Job, error) {
	job := &models.Job{}
	result := mgm.Coll(job).FindOne(mgm.Ctx(), query, opts...)

	if err := result.Err(); err != nil {
		return nil, err
	}
	err := result.Decode(job)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (j *jobService) DeleteJob(id string) error {
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result := mgm.Coll(&models.Job{}).FindOneAndDelete(mgm.Ctx(), bson.M{"_id": pid})
	if err := result.Err(); err != nil {
		return err
	}
	return nil
}
