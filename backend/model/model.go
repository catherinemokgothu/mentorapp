package model

import (
	"com.cathy.mentor-backend/helper"
	"com.cathy.mentor-backend/interfaces"
)

func createUsers() {
	db := helper.ConnectCB()

	users := [2]interfaces.User{
		{Username: "Joe", Email: "joe@gmail.com", Role: "Mentor"},
		{Username: "John", Email: "john.doe.gmail.com", Role: "Mentor"},
	}

	for i := 0; i < len(users); i++ {
		generatePassword := helper.GenerateEncryptPassword([]byte(users[i].Username))
		user := interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatePassword, Role: users[i].Role}
		db.Create(&user)

		mentor := interfaces.Mentor{FirstName: string(users[i].Username), LastName: "Doe", City: "Brighton", Country: "UK", ContactNumber: "1234567", LevelTechnicalExpirience: "Beginner", CurrentTechLevel: "Beginner", SkillsOffersMentoringDesc: "Technical Java Programming & Soft Skills", UserId: user.ID}
		db.Create(&mentor)
	}
	defer db.Close()

}
