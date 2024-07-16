package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/quotation_product_material_sheets"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createQuotationProductMaterialSheetsHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		quotation_product_material_sheetsModel := models.QuotationProductMaterialSheets{}
		c.Bind(&quotation_product_material_sheetsModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newQuotationProductMaterialSheets := models.NewQuotationProductMaterialSheetsRepository(db.DB)
		newQuotationProductMaterialSheets.Create(&quotation_product_material_sheetsModel)

		setFlashmessages(c, "success", "QuotationProductMaterialSheets created successfully!!")
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
	return renderView(c, quotation_product_material_sheets.QuotationProductMaterialSheetsIndex(
		"| Create QuotationProductMaterialSheets",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		quotation_product_material_sheets.CreateQuotationProductMaterialSheets(),
	))

}

func quotation_product_material_sheetsListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newQuotationProductMaterialSheets := models.NewQuotationProductMaterialSheetsRepository(db.DB)
	allQuotationProductMaterialSheets, err := newQuotationProductMaterialSheets.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's QuotationProductMaterialSheets List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, quotation_product_material_sheets.QuotationProductMaterialSheetsIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		quotation_product_material_sheets.QuotationProductMaterialSheetsList(titlePage, allQuotationProductMaterialSheets),
	))
}

func updateQuotationProductMaterialSheetsHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newQuotationProductMaterialSheets := models.NewQuotationProductMaterialSheetsRepository(db.DB)

	quotation_product_material_sheetsModel, err := newQuotationProductMaterialSheets.GetSingle(idParams)

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
	//quotation_product_material_sheetsModel := models.QuotationProductMaterialSheets{}
	c.Bind(quotation_product_material_sheetsModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newQuotationProductMaterialSheets := models.NewQuotationProductMaterialSheetsRepository(db.DB)
	err = newQuotationProductMaterialSheets.Update(quotation_product_material_sheetsModel)

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

	setFlashmessages(c, "success", "QuotationProductMaterialSheets successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/quotation_product_material_sheets/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, quotation_product_material_sheets.QuotationProductMaterialSheetsIndex(
		fmt.Sprintf("| Edit QuotationProductMaterialSheets #%d", quotation_product_material_sheetsModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		quotation_product_material_sheets.UpdateQuotationProductMaterialSheets(*quotation_product_material_sheetsModel, tz.(string)),
	))

}

func deleteQuotationProductMaterialSheetsHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.QuotationProductMaterialSheets{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newQuotationProductMaterialSheets := models.NewQuotationProductMaterialSheetsRepository(db.DB)
	err = newQuotationProductMaterialSheets.Delete(idParams)
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

	setFlashmessages(c, "success", "QuotationProductMaterialSheets successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/quotation_product_material_sheets/list")
}
