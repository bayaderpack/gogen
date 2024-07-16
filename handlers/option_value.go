package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/option_value"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createOptionValueHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	option_valueModel := models.OptionValue{}
	c.Bind(&option_valueModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newOptionValue := models.NewOptionValueRepository(db.DB)
	newOptionValue.Create(&option_valueModel)

setFlashmessages(c, "success", "OptionValue created successfully!!")
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
	return renderView(c, option_value.OptionValueIndex(
		"| Create OptionValue",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		option_value.CreateOptionValue(),
	))

}

func option_valueListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newOptionValue := models.NewOptionValueRepository(db.DB)
	allOptionValue, err := newOptionValue.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's OptionValue List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, option_value.OptionValueIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		option_value.OptionValueList(titlePage, allOptionValue),
	))
}

func updateOptionValueHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newOptionValue := models.NewOptionValueRepository(db.DB)
	
	option_valueModel , err := newOptionValue.GetSingle(idParams)

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
		//option_valueModel := models.OptionValue{}
	c.Bind(option_valueModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newOptionValue := models.NewOptionValueRepository(db.DB)
	err = newOptionValue.Update(option_valueModel)

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


		setFlashmessages(c, "success", "OptionValue successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/option_value/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, option_value.OptionValueIndex(
		fmt.Sprintf("| Edit OptionValue #%d", option_valueModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		option_value.UpdateOptionValue(*option_valueModel, tz.(string)),
	))

}

func deleteOptionValueHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.OptionValue{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newOptionValue := models.NewOptionValueRepository(db.DB)
	err = newOptionValue.Delete(idParams)
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

	setFlashmessages(c, "success", "OptionValue successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/option_value/list")
}
