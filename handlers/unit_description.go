package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/unit_description"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createUnitDescriptionHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	unit_descriptionModel := models.UnitDescription{}
	c.Bind(&unit_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newUnitDescription := models.NewUnitDescriptionRepository(db.DB)
	newUnitDescription.Create(&unit_descriptionModel)

setFlashmessages(c, "success", "UnitDescription created successfully!!")
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
	return renderView(c, unit_description.UnitDescriptionIndex(
		"| Create UnitDescription",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		unit_description.CreateUnitDescription(),
	))

}

func unit_descriptionListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newUnitDescription := models.NewUnitDescriptionRepository(db.DB)
	allUnitDescription, err := newUnitDescription.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's UnitDescription List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, unit_description.UnitDescriptionIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		unit_description.UnitDescriptionList(titlePage, allUnitDescription),
	))
}

func updateUnitDescriptionHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newUnitDescription := models.NewUnitDescriptionRepository(db.DB)
	
	unit_descriptionModel , err := newUnitDescription.GetSingle(idParams)

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
		//unit_descriptionModel := models.UnitDescription{}
	c.Bind(unit_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newUnitDescription := models.NewUnitDescriptionRepository(db.DB)
	err = newUnitDescription.Update(unit_descriptionModel)

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


		setFlashmessages(c, "success", "UnitDescription successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/unit_description/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, unit_description.UnitDescriptionIndex(
		fmt.Sprintf("| Edit UnitDescription #%d", unit_descriptionModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		unit_description.UpdateUnitDescription(*unit_descriptionModel, tz.(string)),
	))

}

func deleteUnitDescriptionHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.UnitDescription{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newUnitDescription := models.NewUnitDescriptionRepository(db.DB)
	err = newUnitDescription.Delete(idParams)
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

	setFlashmessages(c, "success", "UnitDescription successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/unit_description/list")
}
