package controllers

import (
	"ahmadarif/go-revel/app/models"
	"ahmadarif/go-revel/app/utils"

	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type App struct {
	gormc.TxnController
}

func (c App) Index() revel.Result {
	message := "Belajar revel"
	return c.Render(message)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
}

func (c App) GetAllUser() revel.Result {
	var users = []models.User{}
	c.Txn.Find(&users)
	return c.RenderJSON(users)
}

func (c App) InsertUser(name, email, password string) revel.Result {
	c.Validation.Required(name).Message("Name is required!")
	c.Validation.MinSize(name, 3).Message("Name is not long enough!")
	c.Validation.Required(email).Message("Email is required!")
	c.Validation.Email(email).Message("Email is not valid!")
	c.Validation.Required(password).Message("Password is required!")

	if c.Validation.HasErrors() {
		data := make(map[string]interface{})
		data["message"] = "Validation error."
		data["errors"] = utils.ErrorMapToArray(c.Validation.ErrorMap())
		return c.RenderJSON(data)
	}

	user := models.User{Name: name, Email: email, Active: true}
	user.SetNewPassword(password)
	c.Txn.Create(&user)

	return c.RenderJSON(user)
}
