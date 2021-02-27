package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Application struct {
	mgm.DefaultModel `bson:",inline"`
	StudentID        primitive.ObjectID `json:"-" bson:"studentId"`
	Student          *Student           `json:"student" bson:"student"`
	JobID            primitive.ObjectID `json:"jobId" bson:"jobId"`
	Job              *Job               `json:"job" bson:"job"`
	CompanyID        primitive.ObjectID `json:"companyId" bson:"companyId"`
	Company          *Company           `json:"company" bson:"company"`
}
