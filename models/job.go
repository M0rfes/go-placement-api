package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Job defines the structure for a job.
type Job struct {
	mgm.DefaultModel    `bson:",inline"`
	Title               string             `json:"title" bson:"title"`
	Description         string             `json:"description" bson:"description"`
	Openings            uint               `json:"openings" bson:"openings"`
	Type                string             `json:"type" bson:"type"`
	Location            string             `json:"location" bson:"location"`
	Position            string             `json:"position" bson:"position"`
	LastDayOfSummission primitive.DateTime `json:"lastDayOfSummission" bson:"lastDayOfSummission"`
	Company             primitive.ObjectID `json:"company" bson:"company"`
}
