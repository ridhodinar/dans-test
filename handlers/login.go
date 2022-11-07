package handlers

import (
	"fmt"
	"ridhodinar/dans-tes/models"
	"ridhodinar/dans-tes/restapi/operations/user"
	"ridhodinar/dans-tes/utils"

	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type loginImpl struct {
	dbClient *gorm.DB
}

func NewUserLoginHandler(db *gorm.DB) user.LoginHandler {
	return &loginImpl{
		dbClient: db,
	}
}

func (impl *loginImpl) Handle(params user.LoginParams) middleware.Responder {
	username := params.Login.Username
	userDetail := new(models.User)
	result := impl.dbClient.Where("username = ?", username).First(&userDetail)
	if result.Error != nil {
		fmt.Println(result.Error)
		return user.NewLoginInternalServerError()
	}
	err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(*params.Login.Password))
	if err != nil {
		fmt.Println(err)
		return user.NewLoginNotFound()
	}
	token, err := utils.GenerateJWT(userDetail.Username)
	if err != nil {
		return user.NewLoginInternalServerError().WithPayload("Error defining token")
	}
	return user.NewLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: token})
}
