package interfaces

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Role     string
}

type Mentor struct {
	gorm.Model
	FirstName                 string
	LastName                  string
	City                      string
	Country                   string
	ContactNumber             string
	LevelTechnicalExpirience  string
	CurrentTechLevel          string
	SkillsOffersMentoringDesc string
	UserId                    uint
}

type UserResponse struct {
	ID   uint
	Name string
	Role string
}

type MentorResponse struct {
	ID       uint
	Username string
	Email    string
	Users    []UserResponse
}

type Validation struct {
	Value string
	Valid string
}
