package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Job defines the structure for a job.
type Job struct {
	mgm.DefaultModel    `bson:",inline"`
	CTC                 float64            `json:"ctc" bson:"ctc"`
	Description         string             `json:"description" bson:"description"`
	Openings            uint               `json:"openings" bson:"openings"`
	Type                string             `json:"type" bson:"type"`
	Location            string             `json:"location" bson:"location"`
	Position            string             `json:"position" bson:"position"`
	LastDayOfSummission primitive.DateTime `json:"lastDayOfSummission" bson:"lastDayOfSummission"`
	CompanyID           primitive.ObjectID `json:"-" bson:"companyId"`
	Company             *Company           `json:"company" bson:"company"`
	MinCGPA             float32            `json:"minCGPA" bson:"minCGPA"`
}
