package models

import "gorm.io/gorm"

type Officer struct {
	gorm.Model
	Name     string `json:"name" gorm:"text;not null;default:null`
	Email    string `json:"email" gorm:"text;not null;default:null`
	Password string `json:"password" gorm:"text;not null;default:null`
	Reports  []Report  `gorm:"foreignKey:OfficerID"`
}

type CrewMember struct {
	gorm.Model
	Name    string   `gorm:"type:varchar(100);not null"`
	Reports []Report `gorm:"foreignKey`
}

type Captain struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(100);not null"`
	Email    string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string    `gorm:"type:varchar(255);not null"`
}

type Report struct {
	gorm.Model
	OfficerID     uint      `gorm:"not null"` 
	CrewMemberID  uint      `gorm:"not null"` 
	Content       string    `gorm:"type:text;not null"` 
	//Timestamp     time.Time `gorm:"autoCreateTime"`     // Time-stamped when created
	//LastEditedAt  time.Time // Tracks last edit time
	IsEditable    bool      `gorm:"default:true"` 
}


// Reports []Report `gorm:"foreignKey:CrewMemberID",constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`