// Package contact help manage contacts.
package contact

import "time"

// Contact struct is the data we want to work with, 
// the JSON and BSON tags are for MongoDB BSON Data and JSON Data respectively.
type Contact struct {
	FirstName   string    `json:"firstName" bson:"firstName"`
	LastName    string    `json:"lastName" bson:"lastName"`
	Email       string    `json:"email" bson:"email"`
	PhoneNumber string    `json:"phoneNumber" bson:"phoneNumber"`
	Address     string    `json:"address" bson:"address"`
	Company     string    `json:"company" bson:"company"`
	CreatedOn   time.Time `json:"createdOn" bson:"createdOn"`
}
