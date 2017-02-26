package models

type User struct {
	BaseModel
	EncryptedPassword string `json:"-" gorm:"column:password"`
	PlainPassword     string `json:"plainPassword,omitempty" sql:"-"`
	Email             string `json:"email,omitempty"`
}
