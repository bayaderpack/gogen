package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/setting"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createSettingHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	settingModel := models.Setting{}
	c.Bind(&settingModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newSetting := models.NewSettingRepository(db.DB)
	newSetting.Create(&settingModel)

setFlashmessages(c, "success", "Setting created successfully!!")
	c.JSON(http.StatusOK, gin.H{
			"Code":   200,
		"Status": "Ok",
		"Data":   nil,
	})
	}
	username_key_value, ok  := c.Get(username_key)
	if !ok {
		fmt.Println("Some error")
	}
	return renderView(c, setting.SettingIndex(
		"| Create Setting",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		setting.CreateSetting(),
	))

}

func settingListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newSetting := models.NewSettingRepository(db.DB)
	allSetting, err := newSetting.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's Setting List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, setting.SettingIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		setting.SettingList(titlePage, allSetting),
	))
}

func updateSettingHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newSetting := models.NewSettingRepository(db.DB)
	
	settingModel , err := newSetting.GetSingle(idParams)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithError(http.StatusNotFound, fmt.Errorf(
				"something went wrong: %s",
				err,
			))
		}
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf(
				"something went wrong: %s",
				err,
			))
	}
		//settingModel := models.Setting{}
	c.Bind(settingModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newSetting := models.NewSettingRepository(db.DB)
	err = newSetting.Update(settingModel)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {

			c.AbortWithError(http.StatusNotFound, fmt.Errorf(
				"something went wrong: %s",
				err,
			))
			

		}

		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf(
			"something went wrong: %s",
			err,
		))

	}


		setFlashmessages(c, "success", "Setting successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/setting/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, setting.SettingIndex(
		fmt.Sprintf("| Edit Setting #%d", settingModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		setting.UpdateSetting(*settingModel, tz.(string)),
	))

}

func deleteSettingHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.Setting{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newSetting := models.NewSettingRepository(db.DB)
	err = newSetting.Delete(idParams)
	if err != nil {
		if strings.Contains(err.Error(), "an affected row was expected") {

			c.AbortWithError(http.StatusNotFound, fmt.Errorf(
				"something went wrong: %s",
				err,
			))
			

		}

		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf(
			"something went wrong: %s",
			err,
		))

	}

	setFlashmessages(c, "success", "Setting successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/setting/list")
}
