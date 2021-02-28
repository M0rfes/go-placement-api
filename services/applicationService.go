package services

import (
	"placement/models"

	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApplicationService interface {
	CreateApplication(*models.Application) error
	GetAllApplications() *[]*models.Application
	GetApplicationById(id string) (*models.Application, error)
	GetAllApplicationsForAJob(jobID string) (*models.Application, error)
	GetAllApplicationsForCompany(companyID string) (*models.Application, error)
	GetAllApplicationsForStudent(studentID string) (*models.Application, error)
	UpdateApplication(application *models.Application) error
}

type applicationService struct{}

func NewApplicationService() ApplicationService {
	return &applicationService{}
}

func (a *applicationService) CreateApplication(application *models.Application) error {
	return mgm.Coll(application).Create(application)
}

func (a *applicationService) GetAllApplications() *[]*models.Application {
	applications := &[]*models.Application{}
	ctx := mgm.Ctx()
	result, err := mgm.Coll(&models.Application{}).Aggregate(ctx, bson.A{
		bson.M{
			operator.Lookup: bson.M{
				"from":         mgm.Coll(&models.Company{}).Name(),
				"localField":   "companyId",
				"foreignField": "_id",
				"as":           "company",
			},
		},
		bson.M{
			operator.Lookup: bson.M{
				"from":         mgm.Coll(&models.Job{}).Name(),
				"localField":   "jobId",
				"foreignField": "_id",
				"as":           "job",
			},
		},
		bson.M{
			operator.Lookup: bson.M{
				"from":         mgm.Coll(&models.Student{}).Name(),
				"localField":   "studentId",
				"foreignField": "_id",
				"as":           "student",
			},
		},
		bson.M{
			"$unwind": "$company",
		},
		bson.M{
			"$unwind": "$job",
		},
		bson.M{
			"$unwind": "$student",
		},
		bson.M{
			"$project": bson.M{
				"companyId": false,
				"studentId": false,
				"jobId":     false,
				"company": bson.M{
					"password": false,
				},
				"student": bson.M{
					"password": false,
				},
			},
		},
	})
	if err != nil {
		return applications
	}
	result.All(ctx, applications)
	return applications
}

func (a *applicationService) GetApplicationById(id string) (*models.Application, error) {
	applications := []*models.Application{}
	ctx := mgm.Ctx()
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result, err := mgm.Coll(&models.Application{}).Aggregate(ctx, bson.A{
		bson.M{
			"$match": bson.M{"_id": pid},
		},
		bson.M{
			operator.Lookup: bson.M{
				"from":         mgm.Coll(&models.Company{}).Name(),
				"localField":   "companyId",
				"foreignField": "_id",
				"as":           "company",
			},
		},
		bson.M{
			operator.Lookup: bson.M{
				"from":         mgm.Coll(&models.Job{}).Name(),
				"localField":   "jobId",
				"foreignField": "_id",
				"as":           "job",
			},
		},
		bson.M{
			operator.Lookup: bson.M{
				"from":         mgm.Coll(&models.Student{}).Name(),
				"localField":   "studentId",
				"foreignField": "_id",
				"as":           "student",
			},
		},
		bson.M{
			"$unwind": "$company",
		},
		bson.M{
			"$unwind": "$job",
		},
		bson.M{
			"$unwind": "$student",
		},
		bson.M{
			"$project": bson.M{
				"companyId": false,
				"studentId": false,
				"jobId":     false,
				"company": bson.M{
					"password": false,
				},
				"student": bson.M{
					"password": false,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	result.All(ctx, &applications)
	return applications[0], nil
}

func (a *applicationService) GetAllApplicationsForAJob(jobID string) (*models.Application, error) {
	applications := []*models.Application{}
	ctx := mgm.Ctx()
	pid, err := primitive.ObjectIDFromHex(jobID)
	if err != nil {
		return nil, err
	}
	result, err := mgm.Coll(&models.Application{}).Aggregate(ctx, bson.A{
		bson.M{
			"$match": bson.M{"jobId": pid},
		},
		bson.M{
			operator.Lookup: bson.M{
				"from":         mgm.Coll(&models.Company{}).Name(),
				"localField":   "companyId",
				"foreignField": "_id",
				"as":           "company",
			},
		},
		bson.M{
			operator.Lookup: bson.M{
				"from":         mgm.Coll(&models.Job{}).Name(),
				"localField":   "jobId",
				"foreignField": "_id",
				"as":           "job",
			},
		},
		bson.M{
			operator.Lookup: bson.M{
				"from":         mgm.Coll(&models.Student{}).Name(),
				"localField":   "studentId",
				"foreignField": "_id",
				"as":           "student",
			},
		},
		bson.M{
			"$unwind": "$company",
		},
		bson.M{
			"$unwind": "$job",
		},
		bson.M{
			"$unwind": "$student",
		},
		bson.M{
			"$project": bson.M{
				"companyId": false,
				"studentId": false,
				"jobId":     false,
				"company": bson.M{
					"password": false,
				},
				"student": bson.M{
					"password": false,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	result.All(ctx, &applications)
	return applications[0], nil
}

func (a *applicationService) GetAllApplicationsForCompany(companyID string) (*models.Application, error) {
	applications := []*models.Application{}
	ctx := mgm.Ctx()
	pid, err := primitive.ObjectIDFromHex(companyID)
	if err != nil {
		return nil, err
	}
	result, err := mgm.Coll(&models.Application{}).Aggregate(ctx, bson.A{
		bson.M{
			"$match": bson.M{"companyId": pid},
		},
		bson.M{
			operator.Lookup: bson.M{
				"from":         mgm.Coll(&models.Company{}).Name(),
				"localField":   "companyId",
				"foreignField": "_id",
				"as":           "company",
			},
		},
		bson.M{
			operator.Lookup: bson.M{
				"from":         mgm.Coll(&models.Job{}).Name(),
				"localField":   "jobId",
				"foreignField": "_id",
				"as":           "job",
			},
		},
		bson.M{
			operator.Lookup: bson.M{
				"from":         mgm.Coll(&models.Student{}).Name(),
				"localField":   "studentId",
				"foreignField": "_id",
				"as":           "student",
			},
		},
		bson.M{
			"$unwind": "$company",
		},
		bson.M{
			"$unwind": "$job",
		},
		bson.M{
			"$unwind": "$student",
		},
		bson.M{
			"$project": bson.M{
				"companyId": false,
				"studentId": false,
				"jobId":     false,
				"company": bson.M{
					"password": false,
				},
				"student": bson.M{
					"password": false,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	result.All(ctx, &applications)
	return applications[0], nil
}

func (a *applicationService) GetAllApplicationsForStudent(studentID string) (*models.Application, error) {
	applications := []*models.Application{}
	ctx := mgm.Ctx()
	pid, err := primitive.ObjectIDFromHex(studentID)
	if err != nil {
		return nil, err
	}
	result, err := mgm.Coll(&models.Application{}).Aggregate(ctx, bson.A{
		bson.M{
			"$match": bson.M{"studentId": pid},
		},
		bson.M{
			operator.Lookup: bson.M{
				"from":         mgm.Coll(&models.Company{}).Name(),
				"localField":   "companyId",
				"foreignField": "_id",
				"as":           "company",
			},
		},
		bson.M{
			operator.Lookup: bson.M{
				"from":         mgm.Coll(&models.Job{}).Name(),
				"localField":   "jobId",
				"foreignField": "_id",
				"as":           "job",
			},
		},
		bson.M{
			operator.Lookup: bson.M{
				"from":         mgm.Coll(&models.Student{}).Name(),
				"localField":   "studentId",
				"foreignField": "_id",
				"as":           "student",
			},
		},
		bson.M{
			"$unwind": "$company",
		},
		bson.M{
			"$unwind": "$job",
		},
		bson.M{
			"$unwind": "$student",
		},
		bson.M{
			"$project": bson.M{
				"companyId": false,
				"studentId": false,
				"jobId":     false,
				"company": bson.M{
					"password": false,
				},
				"student": bson.M{
					"password": false,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	result.All(ctx, &applications)
	return applications[0], nil
}

func (a *applicationService) UpdateApplication(application *models.Application) error {
	return mgm.Coll(application).Update(application)
}
