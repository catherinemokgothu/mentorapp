package users

import (
	"time"

	"com.cathy.mentor-backend/helper"
	"com.cathy.mentor-backend/interfaces"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func generateToken(user *interfaces.User) string {

	//sign token
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute ^ 60).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helper.HandleError(err)

	return token
}

func prepareResponse(user *interfaces.User, users []interfaces.UserResponse) map[string]interface{} {
	//Create response
	userResponse := &interfaces.MentorResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Users:    users,
	}
	//Response
	var response = map[string]interface{}{"message": "Success"}
	response["jwt"] = generateToken(user)
	response["data"] = userResponse
	return response
}

func Login(username string, password string) map[string]interface{} {
	valid := helper.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: password, Valid: "password"},
		})

	if valid {
		//Connect to database
		db := helper.ConnectCB()
		user := &interfaces.User{}
		if db.Where("username = ?", username).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}

		//verify password
		passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if passwordErr == bcrypt.ErrMismatchedHashAndPassword && passwordErr != nil {
			return map[string]interface{}{"message": "Wrong password"}
		}

		//Find a user record
		mentors := []interfaces.UserResponse{}
		db.Table("users").Select("id, username, role").Where("id = ?", user.ID).Scan(&user)

		defer db.Close()

		var response = prepareResponse(user, mentors)
		return response

	} else {
		return map[string]interface{}{"message": "incorrect values"}
	}

}

func Register(username string, password string, email string, role string, firstname string, lastname string, city string, country string, contactNumber string, levelTechnicalExpirience string, currentTechLevel string, skillsOffersMentoringDesc string) map[string]interface{} {
	valid := helper.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: password, Valid: "password"},
		})

	if valid {
		//Connect to database
		db := helper.ConnectCB()
		generatePassword := helper.GenerateEncryptPassword([]byte(password))
		user := interfaces.User{Username: username, Email: email, Password: generatePassword, Role: role}
		db.Create(&user)

		mentor := interfaces.Mentor{FirstName: firstname, LastName: lastname, City: city, Country: country, ContactNumber: contactNumber, LevelTechnicalExpirience: levelTechnicalExpirience, CurrentTechLevel: currentTechLevel, SkillsOffersMentoringDesc: skillsOffersMentoringDesc, UserId: user.ID}
		db.Create(&mentor)

		defer db.Close()

		mentors := []interfaces.UserResponse{}
		mentorResponse := interfaces.UserResponse{ID: user.ID, Name: user.Username, Role: user.Role}
		mentors = append(mentors, mentorResponse)
		var response = prepareResponse(&user, mentors)
		return response

	} else {
		return map[string]interface{}{"message": "incorrect values"}
	}

}
