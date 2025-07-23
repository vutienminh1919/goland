package models

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	TimeCreated int    `json:"time_created"`
	TimeUpdated int    `json:"time_updated"`
}
