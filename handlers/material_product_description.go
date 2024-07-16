package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/material_product_description"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createMaterialProductDescriptionHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		material_product_descriptionModel := models.MaterialProductDescription{}
		c.Bind(&material_product_descriptionModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newMaterialProductDescription := models.NewMaterialProductDescriptionRepository(db.DB)
		newMaterialProductDescription.Create(&material_product_descriptionModel)

		setFlashmessages(c, "success", "MaterialProductDescription created successfully!!")
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
	return renderView(c, material_product_description.MaterialProductDescriptionIndex(
		"| Create MaterialProductDescription",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		material_product_description.CreateMaterialProductDescription(),
	))

}

func material_product_descriptionListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newMaterialProductDescription := models.NewMaterialProductDescriptionRepository(db.DB)
	allMaterialProductDescription, err := newMaterialProductDescription.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's MaterialProductDescription List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, material_product_description.MaterialProductDescriptionIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		material_product_description.MaterialProductDescriptionList(titlePage, allMaterialProductDescription),
	))
}

func updateMaterialProductDescriptionHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newMaterialProductDescription := models.NewMaterialProductDescriptionRepository(db.DB)

	material_product_descriptionModel, err := newMaterialProductDescription.GetSingle(idParams)

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
	//material_product_descriptionModel := models.MaterialProductDescription{}
	c.Bind(material_product_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newMaterialProductDescription := models.NewMaterialProductDescriptionRepository(db.DB)
	err = newMaterialProductDescription.Update(material_product_descriptionModel)

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

	setFlashmessages(c, "success", "MaterialProductDescription successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/material_product_description/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, material_product_description.MaterialProductDescriptionIndex(
		fmt.Sprintf("| Edit MaterialProductDescription #%d", material_product_descriptionModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		material_product_description.UpdateMaterialProductDescription(*material_product_descriptionModel, tz.(string)),
	))

}

func deleteMaterialProductDescriptionHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.MaterialProductDescription{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newMaterialProductDescription := models.NewMaterialProductDescriptionRepository(db.DB)
	err = newMaterialProductDescription.Delete(idParams)
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

	setFlashmessages(c, "success", "MaterialProductDescription successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/material_product_description/list")
}
