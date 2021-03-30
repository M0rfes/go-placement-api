package models

import "github.com/kamva/mgm/v3"

type Admin struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string `json:"username" bson:"username"`
	Password         string `json:"password" bson:"password"`
}

type ToggleAproven struct {
	StudentID string `json:"studentId"`
}
type ToggleAprovenCompany struct {
	CompanyID string `json:"companyId"`
}
