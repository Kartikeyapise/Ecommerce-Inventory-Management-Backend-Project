package model

var ValidUserTypes = []string{"USER", "MERCHANT"}

//User contains : name and email where email is the primary key
type User struct {
	Email string `json:"email" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null"`
	Type  string `json:"type" gorm:"not null"`
}
