package models

type Driver struct {
	ID        			uint64    `gorm:"primary_key;auto_increment"`
	Name						string		`gorm:"size:50;not null" json:"name"`
	Email     			string    `gorm:"size:100;not null;unique" json:"email"`
}



