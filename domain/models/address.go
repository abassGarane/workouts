package models

type Address struct {
	HomeCounty  string `json:"home_county" bson:"home_county"`
	HomeTown    string `json:"home_town" bson:"home_town"`
	HomeVillage string `json:"home_village" bson:"home_village"`
}
