package models

import "github.com/kamva/mgm/v3"

// Student model that will be saved in the database
type Student struct {
	mgm.DefaultModel `bson:",inline"`
	FirstName        string `json:"firstName" bson:"firstName"`
	LastName         string `json:"lastName" bson:"lastName"`
	Avatar           string `json:"avatar" bson:"avatar"`
	UINNumber        string `json:"uinNumber" bson:"uinNumber"`
	PhoneNumber      string `json:"phoneNumber" bson:"phoneNumber"`
	Gender           string `json:"gender" bson:"gender"`
	Email            string `json:"email" bson:"email"`
	Department       string `json:"department" bson:"department"`
	Program          string `json:"program" bson:"program" bson:"program"`
	CurrentAddress   string `json:"currentAddress" bson:"currentAddress"`
	HomeAddress      string `json:"homeAddress" bson:"homeAddress"`
	Password         string `json:"password" bson:"password"`
	ConfirmPassword  string `json:"confirmPassword" bson:"-"`
	Resume           string `json:"resume" bson:"resume"`
	Sem1             uint8  `json:"sem1" bson:"sem1"`
	Sem2             uint8  `json:"sem2" bson:"sem2"`
	Sem3             uint8  `json:"sem3" bson:"sem3"`
	Sem4             uint8  `json:"sem4" bson:"sem4"`
	Sem5             uint8  `json:"sem5" bson:"sem5"`
	Sem6             uint8  `json:"sem6" bson:"sem6"`
	Sem7             uint8  `json:"sem7" bson:"sem7"`
	Sem8             uint8  `json:"sem8" bson:"sem8"`
	Approved         bool   `json:"approved" bson:"approved"`
}
