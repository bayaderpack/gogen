package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/material_property"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createMaterialPropertyHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	material_propertyModel := models.MaterialProperty{}
	c.Bind(&material_propertyModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newMaterialProperty := models.NewMaterialPropertyRepository(db.DB)
	newMaterialProperty.Create(&material_propertyModel)

setFlashmessages(c, "success", "MaterialProperty created successfully!!")
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
	return renderView(c, material_property.MaterialPropertyIndex(
		"| Create MaterialProperty",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		material_property.CreateMaterialProperty(),
	))

}

func material_propertyListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newMaterialProperty := models.NewMaterialPropertyRepository(db.DB)
	allMaterialProperty, err := newMaterialProperty.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's MaterialProperty List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, material_property.MaterialPropertyIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		material_property.MaterialPropertyList(titlePage, allMaterialProperty),
	))
}

func updateMaterialPropertyHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newMaterialProperty := models.NewMaterialPropertyRepository(db.DB)
	
	material_propertyModel , err := newMaterialProperty.GetSingle(idParams)

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
		//material_propertyModel := models.MaterialProperty{}
	c.Bind(material_propertyModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newMaterialProperty := models.NewMaterialPropertyRepository(db.DB)
	err = newMaterialProperty.Update(material_propertyModel)

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


		setFlashmessages(c, "success", "MaterialProperty successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/material_property/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, material_property.MaterialPropertyIndex(
		fmt.Sprintf("| Edit MaterialProperty #%d", material_propertyModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		material_property.UpdateMaterialProperty(*material_propertyModel, tz.(string)),
	))

}

func deleteMaterialPropertyHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.MaterialProperty{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newMaterialProperty := models.NewMaterialPropertyRepository(db.DB)
	err = newMaterialProperty.Delete(idParams)
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

	setFlashmessages(c, "success", "MaterialProperty successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/material_property/list")
}
