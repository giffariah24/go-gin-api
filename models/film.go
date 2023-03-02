package models

type Film struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Title    string `gorm:"type:varchar(100)" json:"title"`
	Director string `gorm:"type:varchar(100)" json:"director"`
}
