package dto

type StaycationUserResponse struct {
	Id          string `json:"user_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"postcode"`
	DateOfBirth string `json:"dateOfBirth"`
	Status      string `json:"status"`
}
