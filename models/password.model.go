package models

type PasswordLog struct {
	BaseModel
	Request  string `json:"request" gorm:"column:request"`
	Response int    `json:"response" gorm:"column:response"`
}

func (PasswordLog) TableName() string {
	return "password_logs"
}
