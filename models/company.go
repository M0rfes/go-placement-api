package models

import "github.com/kamva/mgm/v3"

// Company defines the structure for company
type Company struct {
	mgm.DefaultModel   `bson:",inline"`
	Name               string `json:"name" bson:"name"`
	Email              string `json:"email" bson:"email"`
	RegistrationNumber string `json:"registrationNumber" bson:"registrationNumber"`
	GSTNumber          string `json:"gstNumber" bson:"gstNumber"`
	WebSiteURL         string `json:"webSiteURL" bson:"webSiteURL"`
	PhoneNumber        string `json:"phoneNumber" bson:"phoneNumber"`
	Address            string `json:"address" bson:"address"`
	Password           string `json:"password" bson:"password"`
	ConfirmPassword    string `json:"confirmPassword" bson:"-"`
	Logo               string `json:"logo" bson:"logo"`
}
