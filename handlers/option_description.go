package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/option_description"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createOptionDescriptionHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		option_descriptionModel := models.OptionDescription{}
		c.Bind(&option_descriptionModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newOptionDescription := models.NewOptionDescriptionRepository(db.DB)
		newOptionDescription.Create(&option_descriptionModel)

		setFlashmessages(c, "success", "OptionDescription created successfully!!")
		c.JSON(http.StatusOK, gin.H{
			"Code":   200,
			"Status": "Ok",
			"Data":   nil,
		})
	}
	username_key_value, ok := c.Get(username_key)
	if !ok {
		fmt.Println("Some error")
	}
	return renderView(c, option_description.OptionDescriptionIndex(
		"| Create OptionDescription",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		option_description.CreateOptionDescription(),
	))

}

func option_descriptionListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newOptionDescription := models.NewOptionDescriptionRepository(db.DB)
	allOptionDescription, err := newOptionDescription.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's OptionDescription List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, option_description.OptionDescriptionIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		option_description.OptionDescriptionList(titlePage, allOptionDescription),
	))
}

func updateOptionDescriptionHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newOptionDescription := models.NewOptionDescriptionRepository(db.DB)

	option_descriptionModel, err := newOptionDescription.GetSingle(idParams)

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
	//option_descriptionModel := models.OptionDescription{}
	c.Bind(option_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newOptionDescription := models.NewOptionDescriptionRepository(db.DB)
	err = newOptionDescription.Update(option_descriptionModel)

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

	setFlashmessages(c, "success", "OptionDescription successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/option_description/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, option_description.OptionDescriptionIndex(
		fmt.Sprintf("| Edit OptionDescription #%d", option_descriptionModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		option_description.UpdateOptionDescription(*option_descriptionModel, tz.(string)),
	))

}

func deleteOptionDescriptionHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.OptionDescription{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newOptionDescription := models.NewOptionDescriptionRepository(db.DB)
	err = newOptionDescription.Delete(idParams)
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

	setFlashmessages(c, "success", "OptionDescription successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/option_description/list")
}
