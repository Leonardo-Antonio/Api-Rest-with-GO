package model

// Alumn is the struct or table of the bd
type Alumn struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Dni      string `json:"dni"`
	Age      uint8  `json:"age"`
}
