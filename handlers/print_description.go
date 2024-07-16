package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/print_description"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createPrintDescriptionHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		print_descriptionModel := models.PrintDescription{}
		c.Bind(&print_descriptionModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newPrintDescription := models.NewPrintDescriptionRepository(db.DB)
		newPrintDescription.Create(&print_descriptionModel)

		setFlashmessages(c, "success", "PrintDescription created successfully!!")
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
	return renderView(c, print_description.PrintDescriptionIndex(
		"| Create PrintDescription",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		print_description.CreatePrintDescription(),
	))

}

func print_descriptionListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newPrintDescription := models.NewPrintDescriptionRepository(db.DB)
	allPrintDescription, err := newPrintDescription.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's PrintDescription List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, print_description.PrintDescriptionIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		print_description.PrintDescriptionList(titlePage, allPrintDescription),
	))
}

func updatePrintDescriptionHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newPrintDescription := models.NewPrintDescriptionRepository(db.DB)

	print_descriptionModel, err := newPrintDescription.GetSingle(idParams)

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
	//print_descriptionModel := models.PrintDescription{}
	c.Bind(print_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newPrintDescription := models.NewPrintDescriptionRepository(db.DB)
	err = newPrintDescription.Update(print_descriptionModel)

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

	setFlashmessages(c, "success", "PrintDescription successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/print_description/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, print_description.PrintDescriptionIndex(
		fmt.Sprintf("| Edit PrintDescription #%d", print_descriptionModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		print_description.UpdatePrintDescription(*print_descriptionModel, tz.(string)),
	))

}

func deletePrintDescriptionHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.PrintDescription{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newPrintDescription := models.NewPrintDescriptionRepository(db.DB)
	err = newPrintDescription.Delete(idParams)
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

	setFlashmessages(c, "success", "PrintDescription successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/print_description/list")
}
