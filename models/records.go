package models

type Records struct {
	OldImage string
	NewImage string 
	EventId string 
	s_no int `gorm:"primaryKey"`

}
