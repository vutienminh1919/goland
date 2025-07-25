package models

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	TimeCreated int    `json:"time_created"`
	TimeUpdated int    `json:"time_updated"`
}

type Debt struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Code          string `json:"code"`
	TransactionID string `json:"transaction_id"`
	TimeCreated   int64  `json:"time_created"`
	TimeUpdated   int64  `json:"time_updated"`
}

type Menu struct {
	ID           int    `json:"id"`
	ParentId     string `json:"parent_id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	Icon         string `json:"icon"`
	DisplayOrder int    `json:"display_order"`
	Status       int    `json:"status"`
	TimeCreated  int    `json:"time_created"`
	TimeUpdated  int    `json:"time_updated"`
}
