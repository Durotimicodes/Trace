package users

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/durotimicodes/trace-backend/api/database"
	"github.com/durotimicodes/trace-backend/helpers"
	"github.com/durotimicodes/trace-backend/models"
)

func prepareToken(user *models.User) string {
	//Sign in jwt Token
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute ^ 10).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	return token

}

func prepareResponse(user *models.User, account []models.ResponseAccount, withToken bool) map[string]interface{} {
	//Setup response
	responseUser := models.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: account,
	}

	//Prepare response
	var response = map[string]interface{}{"message": "All is fine"}
	if withToken {
		token := prepareToken(user)
		response["jwt"] = token
	}
	response["data"] = responseUser

	return response

}

func Login(username string, pass string) map[string]interface{} {

	//Before connect to db validate credentials
	valid := helpers.Validation(
		[]models.Validation{
			{Value: username, Valid: "username"},
			{Value: pass, Valid: "password"},
		})

	if valid {

		user := &models.User{}
		if database.DB.Where("username = ?", username).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}

		//Find account for the user
		account := []models.ResponseAccount{}
		if database.DB.Table("account").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&account).RecordNotFound() {
			return map[string]interface{}{"message": "User account not found"}
		}

		var response = prepareResponse(user, account, true)

		return response

	} else {
		return map[string]interface{}{"message": "Not valid credentials"}
	}

}

func Register(username, email, pass string) map[string]interface{} {
	valid := helpers.Validation(
		[]models.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})

	if valid {
		//Create registration Logic
		generatePassword := helpers.HashAndSalt([]byte(pass))
		user := models.User{
			Username: username,
			Email:    email,
			Password: generatePassword,
		}
		database.DB.Create(&user)

		account := models.Account{
			Type:    "Savings Account",
			Name:    string(username + "'s" + " account"),
			Balance: uint(0),
			UserID:  user.ID,
		}
		database.DB.Create(&account)

		accounts := []models.ResponseAccount{}
		respAccount := models.ResponseAccount{
			ID:      account.ID,
			Name:    account.Name,
			Balance: int(account.Balance),
		}
		accounts = append(accounts, respAccount)

		var response = prepareResponse(&user, accounts, true)

		return response

	} else {
		return map[string]interface{}{"message": "Not valid credentials"}
	}
}

func GetUser(id string, jwt string) map[string]interface{} {
	isValid := helpers.ValidateToken(id, jwt)

	//Find and return User
	if isValid {

		user := &models.User{}
		if database.DB.Where("id = ?", id).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}

		//Find account for the user
		account := []models.ResponseAccount{}
		database.DB.Table("account").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&account)

		var response = prepareResponse(user, account, false)
		return response

	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}

}
