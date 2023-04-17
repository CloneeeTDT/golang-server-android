package models

type Word struct {
	ID         uint `gorm:"primary_key;auto_increment;not_null"`
	Word       string
	WordType   string
	Definition string
}
