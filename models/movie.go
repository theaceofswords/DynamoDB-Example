package models

type Movie struct {
	Year      int   `json:"Year"`
	Title     string `json:"Title"`
	Category  string  `json:"Category"`
	Plot      string  `json:"Plot"`
	Rating    float64 `json:"Rating"`
	Director  string  `json:"Director"`
	LeadActor string  `json:"LeadActor"`
	Duration  int   `json:"Duration"`
}


