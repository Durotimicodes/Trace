package users

import (
	"time"

	"github.com/dgrijalva/jwt-go"
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
		//Connect db
		db := helpers.ConnectDB()

		user := &models.User{}
		if db.Where("username = ?", username).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}

		//Find account for the user
		account := []models.ResponseAccount{}
		db.Table("account").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&account)
		defer db.Close()

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
		db := helpers.ConnectDB()
		generatePassword := helpers.HashAndSalt([]byte(pass))
		user := models.User{
			Username: username,
			Email:    email,
			Password: generatePassword,
		}
		db.Create(&user)

		account := models.Account{
			Type:    "Savings Account",
			Name:    string(username + "'s" + " account"),
			Balance: uint(0),
			UserID:  user.ID,
		}
		db.Create(&account)

		defer db.Close()

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

	if isValid {
		db := helpers.ConnectDB()

		user := &models.User{}
		if db.Where("id = ?", id).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}

		//Find account for the user
		account := []models.ResponseAccount{}
		db.Table("account").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&account)
		defer db.Close()

		var response = prepareResponse(user, account, false)
		return response

	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}

}
