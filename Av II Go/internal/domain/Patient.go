package domain

type Patient struct {
	Id           int    `json:"id"`
	Name         string `json:"name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	Rg           string `json:"rg" binding:"required"`
	RegisterDate string `json:"register_date" binding:"required"`
}