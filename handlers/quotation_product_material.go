package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/quotation_product_material"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createQuotationProductMaterialHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	quotation_product_materialModel := models.QuotationProductMaterial{}
	c.Bind(&quotation_product_materialModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newQuotationProductMaterial := models.NewQuotationProductMaterialRepository(db.DB)
	newQuotationProductMaterial.Create(&quotation_product_materialModel)

setFlashmessages(c, "success", "QuotationProductMaterial created successfully!!")
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
	return renderView(c, quotation_product_material.QuotationProductMaterialIndex(
		"| Create QuotationProductMaterial",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		quotation_product_material.CreateQuotationProductMaterial(),
	))

}

func quotation_product_materialListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newQuotationProductMaterial := models.NewQuotationProductMaterialRepository(db.DB)
	allQuotationProductMaterial, err := newQuotationProductMaterial.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's QuotationProductMaterial List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, quotation_product_material.QuotationProductMaterialIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		quotation_product_material.QuotationProductMaterialList(titlePage, allQuotationProductMaterial),
	))
}

func updateQuotationProductMaterialHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newQuotationProductMaterial := models.NewQuotationProductMaterialRepository(db.DB)
	
	quotation_product_materialModel , err := newQuotationProductMaterial.GetSingle(idParams)

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
		//quotation_product_materialModel := models.QuotationProductMaterial{}
	c.Bind(quotation_product_materialModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newQuotationProductMaterial := models.NewQuotationProductMaterialRepository(db.DB)
	err = newQuotationProductMaterial.Update(quotation_product_materialModel)

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


		setFlashmessages(c, "success", "QuotationProductMaterial successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/quotation_product_material/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, quotation_product_material.QuotationProductMaterialIndex(
		fmt.Sprintf("| Edit QuotationProductMaterial #%d", quotation_product_materialModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		quotation_product_material.UpdateQuotationProductMaterial(*quotation_product_materialModel, tz.(string)),
	))

}

func deleteQuotationProductMaterialHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.QuotationProductMaterial{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newQuotationProductMaterial := models.NewQuotationProductMaterialRepository(db.DB)
	err = newQuotationProductMaterial.Delete(idParams)
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

	setFlashmessages(c, "success", "QuotationProductMaterial successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/quotation_product_material/list")
}
