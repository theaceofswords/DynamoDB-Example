package models

type Records struct {
	OldImage string  `json:"OldImage"`
	NewImage string `json:"NewImage"`
	EventId string   `json:"EventId"`
	EventName string  `json:"EventName"`
	s_no int `gorm:"primaryKey"`

}

// type NewImageStruct struct {
//     Year      int
// 	Title     string
// 	Category  string
// 	Plot      string
// 	Rating    float64
// 	Director  string
// 	LeadActor string
// 	Duration  int
// }

// type OldImageStruct struct {
//     oldImage string `json:"oldImage"`
// }
