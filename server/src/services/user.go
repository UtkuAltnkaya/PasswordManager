package services

//TODO Support for multiple accounts
import (
	"encoding/hex"
	"fmt"
	"github/exmaximum/passwordmanager/src/hash"
	"github/exmaximum/passwordmanager/src/helpers"
	"github/exmaximum/passwordmanager/src/models"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserConnection struct {
	*gorm.DB
}

type UserService interface {
	CreateUser(*models.User) helpers.Message[string]
	SignUser(*models.User) helpers.Message[string]
	AddPassword(string, string, int) helpers.Message[interface{}]
	UpdateSitePassword(string, string, int) helpers.Message[interface{}]
	DeleteSitePassword(string, string) helpers.Message[string]
	GetUserPassword(string, string) helpers.Message[interface{}]
	GetUserInfo(string) helpers.Message[interface{}]
	getPasswordName(string) models.PasswordNames
}

func NewUserService(db *gorm.DB) UserConnection {
	return UserConnection{db}
}

func (db UserConnection) CreateUser(user *models.User) helpers.Message[string] {
	hashPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return helpers.NewMessage(err.Error(), false)
	}
	user.Password = hashPassword
	tx := db.Create(user)
	if tx.Error != nil {
		return helpers.NewMessage("User already exist", false)
	}
	return helpers.NewMessage("User created", true)

}
func (db UserConnection) SignUser(user *models.User) helpers.Message[string] {
	var dbUser models.User
	db.Where("email = ?", user.Email).First(&dbUser)
	if len(dbUser.Email) == 0 || !helpers.CheckPassword(dbUser.Password, user.Password) {
		return helpers.NewMessage("Wrong credentials", false)
	}
	*user = dbUser
	return helpers.NewMessage("Logged in", true)
}

func (db UserConnection) AddPassword(userId, site string, length int) helpers.Message[interface{}] {
	password := hash.NewPassword(length)
	encrypted, err := password.Encrypt([]byte(os.Getenv("SALT")))
	if err != nil {
		return helpers.NewMessage[interface{}](err.Error(), false)
	}
	passwordNames := db.getPasswordName(site)
	if len(passwordNames.Name) == 0 {
		passwordNames.Name = site
		tx := db.Create(&passwordNames)
		if tx.Error != nil {
			return helpers.NewMessage[interface{}](tx.Error.Error(), false)
		}
	}
	var item models.Passwords
	db.Where("password_name_id = ? AND user_id = ?", passwordNames.ID, userId).First(&item)
	if len(item.Password) != 0 {
		return helpers.NewMessage[interface{}](fmt.Sprintf("%v is already exist", site), false)
	}
	item = models.Passwords{
		PasswordNameId: passwordNames.ID,
		Password:       hex.EncodeToString(encrypted),
		UserId:         uuid.MustParse(userId),
	}
	tx := db.Create(&item)
	if tx.Error != nil {
		return helpers.NewMessage[interface{}](tx.Error.Error(), false)
	}
	item.PasswordName = passwordNames
	item.Password = string(password.Password)
	return helpers.NewMessage[interface{}](item, true)
}

func (db UserConnection) UpdateSitePassword(userId, site string, length int) helpers.Message[interface{}] {
	passwordNames := db.getPasswordName(site)
	if len(passwordNames.Name) == 0 {
		return helpers.NewMessage[interface{}]("Site not found", false)
	}
	password := hash.NewPassword(length)
	encrypted, err := password.Encrypt([]byte(os.Getenv("SALT")))
	if err != nil {
		return helpers.NewMessage[interface{}](err.Error(), false)
	}
	item := models.Passwords{
		Password: hex.EncodeToString(encrypted),
	}

	tx := db.Model(&item).Clauses(clause.Returning{}).Where("password_name_id = ? AND user_id = ?", passwordNames.ID, userId).Update("Password", item.Password)
	if tx.Error != nil {
		return helpers.NewMessage[interface{}](tx.Error.Error(), false)
	}

	item.Password = string(password.Password)
	item.PasswordName = passwordNames
	return helpers.NewMessage[interface{}](item, true)
}

func (db UserConnection) DeleteSitePassword(userId, site string) helpers.Message[string] {
	passwordNames := db.getPasswordName(site)
	if len(passwordNames.Name) == 0 {
		return helpers.NewMessage("Site not found", false)
	}
	password := models.Passwords{
		PasswordNameId: passwordNames.ID,
		UserId:         uuid.MustParse(userId),
	}
	tx := db.Where("password_name_id = ? AND user_id = ?", passwordNames.ID, userId).Delete(&password)
	if tx.Error != nil {
		return helpers.NewMessage(tx.Error.Error(), false)
	}
	return helpers.NewMessage("Password deleted", true)
}

func (db UserConnection) GetUserPassword(query, userId string) helpers.Message[interface{}] {
	var pass models.Passwords
	passNames := db.getPasswordName(query)
	if len(passNames.Name) == 0 {
		return helpers.NewMessage[interface{}]("Site not found", false)
	}
	db.Where("password_name_id = ? AND user_id = ?", passNames.ID, userId).First(&pass)
	if len(pass.Password) == 0 {
		return helpers.NewMessage[interface{}]("Site not found", false)
	}
	decodeByte, err := hex.DecodeString(pass.Password)
	if err != nil {
		return helpers.NewMessage[interface{}](err.Error(), false)
	}
	decrypted, err := hash.Decrypt([]byte(os.Getenv("SALT")), decodeByte)
	if err != nil {
		return helpers.NewMessage[interface{}](err.Error(), false)
	}
	pass.Password = string(decrypted)
	pass.PasswordName = passNames
	return helpers.NewMessage[interface{}](pass, true)
}

func (db UserConnection) GetUserInfo(userId string) helpers.Message[interface{}] {
	var user models.User
	tx := db.Where("id = ?", userId).Preload("PasswordList.PasswordName").Preload("PasswordList", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"id", "password_name_id", "user_id"})
	}).Find(&user)

	if tx.Error != nil {
		return helpers.NewMessage[interface{}](tx.Error.Error(), false)
	}
	user.Password = ""
	return helpers.NewMessage[interface{}](user, true)
}

func (db UserConnection) getPasswordName(site string) models.PasswordNames {
	var passwordNames models.PasswordNames
	db.Where("Lower(name) = Lower(?)", site).First(&passwordNames)
	return passwordNames
}
