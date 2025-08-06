package amocrm

import "time"

type Account struct {
	AccountID    int       `gorm:"primaryKey"`
	Domain       string    `gorm:"type:varchar(255);unique"`
	AccessToken  string    `gorm:"type:text"`
	RefreshToken string    `gorm:"type:text"`
	Expires      int64     `gorm:"type:int"`
	IssuedAt     time.Time `gorm:"type:datetime"`
	UnisenderKey string    `gorm:"type:text"`
	Contacts     []Contact `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE"`
}
