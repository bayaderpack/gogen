package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/material_property_description"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createMaterialPropertyDescriptionHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	material_property_descriptionModel := models.MaterialPropertyDescription{}
	c.Bind(&material_property_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newMaterialPropertyDescription := models.NewMaterialPropertyDescriptionRepository(db.DB)
	newMaterialPropertyDescription.Create(&material_property_descriptionModel)

setFlashmessages(c, "success", "MaterialPropertyDescription created successfully!!")
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
	return renderView(c, material_property_description.MaterialPropertyDescriptionIndex(
		"| Create MaterialPropertyDescription",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		material_property_description.CreateMaterialPropertyDescription(),
	))

}

func material_property_descriptionListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newMaterialPropertyDescription := models.NewMaterialPropertyDescriptionRepository(db.DB)
	allMaterialPropertyDescription, err := newMaterialPropertyDescription.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's MaterialPropertyDescription List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, material_property_description.MaterialPropertyDescriptionIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		material_property_description.MaterialPropertyDescriptionList(titlePage, allMaterialPropertyDescription),
	))
}

func updateMaterialPropertyDescriptionHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newMaterialPropertyDescription := models.NewMaterialPropertyDescriptionRepository(db.DB)
	
	material_property_descriptionModel , err := newMaterialPropertyDescription.GetSingle(idParams)

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
		//material_property_descriptionModel := models.MaterialPropertyDescription{}
	c.Bind(material_property_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newMaterialPropertyDescription := models.NewMaterialPropertyDescriptionRepository(db.DB)
	err = newMaterialPropertyDescription.Update(material_property_descriptionModel)

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


		setFlashmessages(c, "success", "MaterialPropertyDescription successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/material_property_description/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, material_property_description.MaterialPropertyDescriptionIndex(
		fmt.Sprintf("| Edit MaterialPropertyDescription #%d", material_property_descriptionModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		material_property_description.UpdateMaterialPropertyDescription(*material_property_descriptionModel, tz.(string)),
	))

}

func deleteMaterialPropertyDescriptionHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.MaterialPropertyDescription{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newMaterialPropertyDescription := models.NewMaterialPropertyDescriptionRepository(db.DB)
	err = newMaterialPropertyDescription.Delete(idParams)
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

	setFlashmessages(c, "success", "MaterialPropertyDescription successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/material_property_description/list")
}
