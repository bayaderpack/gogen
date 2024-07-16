package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/sheet_description"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createSheetDescriptionHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		sheet_descriptionModel := models.SheetDescription{}
		c.Bind(&sheet_descriptionModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newSheetDescription := models.NewSheetDescriptionRepository(db.DB)
		newSheetDescription.Create(&sheet_descriptionModel)

		setFlashmessages(c, "success", "SheetDescription created successfully!!")
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
	return renderView(c, sheet_description.SheetDescriptionIndex(
		"| Create SheetDescription",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		sheet_description.CreateSheetDescription(),
	))

}

func sheet_descriptionListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newSheetDescription := models.NewSheetDescriptionRepository(db.DB)
	allSheetDescription, err := newSheetDescription.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's SheetDescription List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, sheet_description.SheetDescriptionIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		sheet_description.SheetDescriptionList(titlePage, allSheetDescription),
	))
}

func updateSheetDescriptionHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newSheetDescription := models.NewSheetDescriptionRepository(db.DB)

	sheet_descriptionModel, err := newSheetDescription.GetSingle(idParams)

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
	//sheet_descriptionModel := models.SheetDescription{}
	c.Bind(sheet_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newSheetDescription := models.NewSheetDescriptionRepository(db.DB)
	err = newSheetDescription.Update(sheet_descriptionModel)

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

	setFlashmessages(c, "success", "SheetDescription successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/sheet_description/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, sheet_description.SheetDescriptionIndex(
		fmt.Sprintf("| Edit SheetDescription #%d", sheet_descriptionModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		sheet_description.UpdateSheetDescription(*sheet_descriptionModel, tz.(string)),
	))

}

func deleteSheetDescriptionHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.SheetDescription{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newSheetDescription := models.NewSheetDescriptionRepository(db.DB)
	err = newSheetDescription.Delete(idParams)
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

	setFlashmessages(c, "success", "SheetDescription successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/sheet_description/list")
}
