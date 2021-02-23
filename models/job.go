package models

import "github.com/kamva/mgm/v3"

// Job defines the structure for a job.
type Job struct {
	mgm.DefaultModel    `bson:",inline"`
	Title               string `json:"title" bson:"title"`
	Description         string `json:"description" bson:"description"`
	Openings            uint   `json:"openings" bson:"openings"`
	Type                string `json:"type" bson:"type"`
	Locals              string `json:"locals" bson:"locals"`
	Position            string `json:"position" bson:"position"`
	LastDayOfSummission string `json:"lastDayOfSummission" bson:"lastDayOfSummission"`
	Company             string `json:"company" bson:"company"`
}
