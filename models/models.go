package models

// User struct
type User struct {
	ID      int     `json:"id"`
	Fname   string  `json:"fname"`
	City    string  `json:"city"`
	Phone   uint64  `json:"phone"`
	Height  float64 `json:"height"`
	Married bool    `json:"married"`
}
