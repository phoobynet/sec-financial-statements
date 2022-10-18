package sics

type SIC struct {
	Code          string `gorm:"primaryKey"`
	IndustryTitle string
	Office        string
}
