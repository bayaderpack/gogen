package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/option_value_description"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createOptionValueDescriptionHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		option_value_descriptionModel := models.OptionValueDescription{}
		c.Bind(&option_value_descriptionModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newOptionValueDescription := models.NewOptionValueDescriptionRepository(db.DB)
		newOptionValueDescription.Create(&option_value_descriptionModel)

		setFlashmessages(c, "success", "OptionValueDescription created successfully!!")
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
	return renderView(c, option_value_description.OptionValueDescriptionIndex(
		"| Create OptionValueDescription",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		option_value_description.CreateOptionValueDescription(),
	))

}

func option_value_descriptionListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newOptionValueDescription := models.NewOptionValueDescriptionRepository(db.DB)
	allOptionValueDescription, err := newOptionValueDescription.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's OptionValueDescription List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, option_value_description.OptionValueDescriptionIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		option_value_description.OptionValueDescriptionList(titlePage, allOptionValueDescription),
	))
}

func updateOptionValueDescriptionHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newOptionValueDescription := models.NewOptionValueDescriptionRepository(db.DB)

	option_value_descriptionModel, err := newOptionValueDescription.GetSingle(idParams)

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
	//option_value_descriptionModel := models.OptionValueDescription{}
	c.Bind(option_value_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newOptionValueDescription := models.NewOptionValueDescriptionRepository(db.DB)
	err = newOptionValueDescription.Update(option_value_descriptionModel)

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

	setFlashmessages(c, "success", "OptionValueDescription successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/option_value_description/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, option_value_description.OptionValueDescriptionIndex(
		fmt.Sprintf("| Edit OptionValueDescription #%d", option_value_descriptionModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		option_value_description.UpdateOptionValueDescription(*option_value_descriptionModel, tz.(string)),
	))

}

func deleteOptionValueDescriptionHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.OptionValueDescription{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newOptionValueDescription := models.NewOptionValueDescriptionRepository(db.DB)
	err = newOptionValueDescription.Delete(idParams)
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

	setFlashmessages(c, "success", "OptionValueDescription successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/option_value_description/list")
}
