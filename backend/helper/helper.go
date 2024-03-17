package helper

import (
	"regexp"

	"com.cathy.mentor-backend/interfaces"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

// function to hadle errors
func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func GenerateEncryptPassword(password []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	HandleError(err)
	return string(hashed)
}

func ConnectCB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=password sslmode=disable")
	HandleError(err)
	return db
}

func Validation(values []interfaces.Validation) bool {
	username := regexp.MustCompile(`([A-Za-z0-9]{5,})+$`)
	email := regexp.MustCompile(`[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z0-9]+$`)

	for i := 0; i < len(values); i++ {
		switch values[i].Valid {
		case "username":
			if !username.MatchString(values[i].Value) {
				return false
			}
		case "email":
			if !email.MatchString(values[i].Value) {
				return false
			}
		case "password":
			if len(values[i].Value) < 5 {
				return false
			}
		}
	}
	return true
}
