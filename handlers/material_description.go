package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/material_description"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createMaterialDescriptionHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	material_descriptionModel := models.MaterialDescription{}
	c.Bind(&material_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newMaterialDescription := models.NewMaterialDescriptionRepository(db.DB)
	newMaterialDescription.Create(&material_descriptionModel)

setFlashmessages(c, "success", "MaterialDescription created successfully!!")
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
	return renderView(c, material_description.MaterialDescriptionIndex(
		"| Create MaterialDescription",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		material_description.CreateMaterialDescription(),
	))

}

func material_descriptionListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newMaterialDescription := models.NewMaterialDescriptionRepository(db.DB)
	allMaterialDescription, err := newMaterialDescription.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's MaterialDescription List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, material_description.MaterialDescriptionIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		material_description.MaterialDescriptionList(titlePage, allMaterialDescription),
	))
}

func updateMaterialDescriptionHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newMaterialDescription := models.NewMaterialDescriptionRepository(db.DB)
	
	material_descriptionModel , err := newMaterialDescription.GetSingle(idParams)

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
		//material_descriptionModel := models.MaterialDescription{}
	c.Bind(material_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newMaterialDescription := models.NewMaterialDescriptionRepository(db.DB)
	err = newMaterialDescription.Update(material_descriptionModel)

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


		setFlashmessages(c, "success", "MaterialDescription successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/material_description/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, material_description.MaterialDescriptionIndex(
		fmt.Sprintf("| Edit MaterialDescription #%d", material_descriptionModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		material_description.UpdateMaterialDescription(*material_descriptionModel, tz.(string)),
	))

}

func deleteMaterialDescriptionHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.MaterialDescription{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newMaterialDescription := models.NewMaterialDescriptionRepository(db.DB)
	err = newMaterialDescription.Delete(idParams)
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

	setFlashmessages(c, "success", "MaterialDescription successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/material_description/list")
}
