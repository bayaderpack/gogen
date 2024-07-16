package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/material_product"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createMaterialProductHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		material_productModel := models.MaterialProduct{}
		c.Bind(&material_productModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newMaterialProduct := models.NewMaterialProductRepository(db.DB)
		newMaterialProduct.Create(&material_productModel)

		setFlashmessages(c, "success", "MaterialProduct created successfully!!")
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
	return renderView(c, material_product.MaterialProductIndex(
		"| Create MaterialProduct",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		material_product.CreateMaterialProduct(),
	))

}

func material_productListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newMaterialProduct := models.NewMaterialProductRepository(db.DB)
	allMaterialProduct, err := newMaterialProduct.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's MaterialProduct List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, material_product.MaterialProductIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		material_product.MaterialProductList(titlePage, allMaterialProduct),
	))
}

func updateMaterialProductHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newMaterialProduct := models.NewMaterialProductRepository(db.DB)

	material_productModel, err := newMaterialProduct.GetSingle(idParams)

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
	//material_productModel := models.MaterialProduct{}
	c.Bind(material_productModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newMaterialProduct := models.NewMaterialProductRepository(db.DB)
	err = newMaterialProduct.Update(material_productModel)

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

	setFlashmessages(c, "success", "MaterialProduct successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/material_product/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, material_product.MaterialProductIndex(
		fmt.Sprintf("| Edit MaterialProduct #%d", material_productModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		material_product.UpdateMaterialProduct(*material_productModel, tz.(string)),
	))

}

func deleteMaterialProductHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.MaterialProduct{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newMaterialProduct := models.NewMaterialProductRepository(db.DB)
	err = newMaterialProduct.Delete(idParams)
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

	setFlashmessages(c, "success", "MaterialProduct successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/material_product/list")
}
