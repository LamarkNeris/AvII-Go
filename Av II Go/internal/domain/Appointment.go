package domain

type Appointment struct {
	Id          int     `json:"id"`
	Patient     Patient `json:"patient"`
	Dentist     Dentist `json:"dentist"`
	DateHour        string  `json:"date_hour" binding:"required"`
	Description string  `json:"description" binding:"required"`
}