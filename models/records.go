package models

type Records struct {
	OldImage string
	NewImage string 
	EventId string 
	EventName string
	s_no int `gorm:"primaryKey"`

}
