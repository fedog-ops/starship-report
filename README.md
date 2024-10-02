# Starship Crew Development Log 🚀


## Overview
The Starship Crew Development Log is a tool designed to help starship officers track the performance, behavior, and development of crew members during long-term space missions. Officers can log and review crew observations to ensure peak performance and well-being throughout deep space exploration.

This system allows officers to maintain and review detailed crew reports, while the ship’s captain has full oversight of crew development. Access to these records is tightly controlled to maintain the security of mission-critical data.

## Features
1. Officer Role

Enter Crew Reports: Officers can log text-based reports about crew members. These logs are time-stamped for historical accuracy.

Multiple Entries: Officers can make multiple observations about a crew member over time, recording development or concerns during the mission.

Editable within One Month: Officers have a one-month window to view and edit their own reports. After this time, logs are locked for editing to preserve the integrity of the records.

2. Captain Role

Full Observation Overview: The Captain can view a time-ordered history of all observations made about a crew member.

Generate Full Reports: The Captain can generate a complete report of all crew members' development logs, which can be downloaded as mission-critical data.

Manage Users: The Captain can assign officers (who log observations) and maintain the crew roster. Non-assigned personnel cannot access the system or its data.

Revoke Officer Access: The Captain can revoke access to an officer at any time, ensuring control over who can log or view crew performance reports.
Roles and Permissions

3. Officers:

Can log reports about crew members.
Can view and edit their own reports within a one-month window.
Captain:

Can view all crew reports and download full mission logs.
Can assign officers and manage crew access.
Has the ability to revoke officer access to the system.
Crew Members:

No direct access to the system or underlying data.


------------------------------------
```go
package models

import (
	"time"
	"gorm.io/gorm"
)

// BaseModel provides common fields for all tables.
type BaseModel struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Officer represents the officer role, which can log and edit crew reports.
type Officer struct {
	BaseModel
	Name     string    `gorm:"type:varchar(100);not null"`
	Email    string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string    `gorm:"type:varchar(255);not null"` // for hashed passwords
	Reports  []Report  `gorm:"foreignKey:OfficerID"`
}

// Captain represents the captain role, with elevated privileges to view all reports and manage users.
type Captain struct {
	BaseModel
	Name     string    `gorm:"type:varchar(100);not null"`
	Email    string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string    `gorm:"type:varchar(255);not null"` // for hashed passwords
	// Can assign officers and revoke their access (this logic would be implemented in services)
}

// CrewMember represents a crew member who is being observed by officers.
type CrewMember struct {
	BaseModel
	Name    string   `gorm:"type:varchar(100);not null"`
	Reports []Report `gorm:"foreignKey:CrewMemberID"`
}

// Report represents the text-based logs entered by officers about a crew member.
type Report struct {
	BaseModel
	OfficerID     uint      `gorm:"not null"` // Foreign key to the officer who created the report
	CrewMemberID  uint      `gorm:"not null"` // Foreign key to the crew member the report is about
	Content       string    `gorm:"type:text;not null"` // Report content
	Timestamp     time.Time `gorm:"autoCreateTime"`     // Time-stamped when created
	LastEditedAt  time.Time // Tracks last edit time
	IsEditable    bool      `gorm:"default:true"`       // Determines if the report is still editable
}

// OfficerAccessRevoked is a table to log officer access revocation actions performed by the captain.
type OfficerAccessRevoked struct {
	BaseModel
	CaptainID uint      `gorm:"not null"` // Foreign key to the captain who revoked access
	OfficerID uint      `gorm:"not null"` // Foreign key to the officer whose access was revoked
	RevokedAt time.Time `gorm:"autoCreateTime"` // Time-stamped when the access was revoked
}

// UserAssignment represents the assignment of officers to crew members (captains handle assignments).
type UserAssignment struct {
	BaseModel
	CaptainID   uint      `gorm:"not null"` // Foreign key to the captain who made the assignment
	OfficerID   uint      `gorm:"not null"` // Foreign key to the officer being assigned
	CrewMemberID uint     `gorm:"not null"` // Foreign key to the crew member being observed
	AssignedAt  time.Time `gorm:"autoCreateTime"` // Time-stamped when the assignment was made
}

// Create function to automatically run migrations for all models.
func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(
		&Officer{},
		&Captain{},
		&CrewMember{},
		&Report{},
		&OfficerAccessRevoked{},
		&UserAssignment{},
	)
	return err
}
```
```JSON

Officers

{
  "name": "Indy",
  "email": "ijones@marshallcollege.edu",
  "password": "ihatesnakes"
}


Crew

{
  "name": "First crew member"
}


Report

{
 	"OfficerID": 1, 
	"CrewMemberID": 1, 
	"Content": "You are not very good at this"
} 

Captain

{
			Name:     "Captain Vader",
			Email:    "vader@starship.com",
			Password: "darkside",
		}
```


root
|
-/backend
|	|
|	/cmd
|	|	|
|	|	main.go
|	Dockerfile
docker-compose.yml