package models

import "github.com/kamva/mgm/v3"

// Student model that will be saved in the database
type Student struct {
	mgm.DefaultModel `bson:",inline"`
	FirstName        string `json:"firstName" bson:"firstName"`
	LastName         string `json:"lastName" bson:"lastName"`
	Avatar           string `json:"avatar" bson:"avatar"`
	UINNumber        string `json:"UINNumber" bson:"UINNumber"`
	PhoneNumber      string `json:"PhoneNumber" bson:"phoneNumber"`
	Gender           string `json:"Gender" bson:"gender"`
	Email            string `json:"email" bson:"email"`
	Department       string `json:"department" bson:"department"`
	Program          string `json:"program" bson:"program" bson:"program"`
	CurrentAddress   string `json:"currentAddress" bson:"currentAddress"`
	HomeAddress      string `json:"homeAddress" bson:"homeAddress"`
	Password         string `json:"password" bson:"password"`
	ConfirmPassword  string `json:"confirmPassword" bson:"-"`
}
