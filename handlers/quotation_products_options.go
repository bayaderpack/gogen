package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/quotation_products_options"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createQuotationProductsOptionsHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		quotation_products_optionsModel := models.QuotationProductsOptions{}
		c.Bind(&quotation_products_optionsModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newQuotationProductsOptions := models.NewQuotationProductsOptionsRepository(db.DB)
		newQuotationProductsOptions.Create(&quotation_products_optionsModel)

		setFlashmessages(c, "success", "QuotationProductsOptions created successfully!!")
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
	return renderView(c, quotation_products_options.QuotationProductsOptionsIndex(
		"| Create QuotationProductsOptions",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		quotation_products_options.CreateQuotationProductsOptions(),
	))

}

func quotation_products_optionsListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newQuotationProductsOptions := models.NewQuotationProductsOptionsRepository(db.DB)
	allQuotationProductsOptions, err := newQuotationProductsOptions.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's QuotationProductsOptions List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, quotation_products_options.QuotationProductsOptionsIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		quotation_products_options.QuotationProductsOptionsList(titlePage, allQuotationProductsOptions),
	))
}

func updateQuotationProductsOptionsHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newQuotationProductsOptions := models.NewQuotationProductsOptionsRepository(db.DB)

	quotation_products_optionsModel, err := newQuotationProductsOptions.GetSingle(idParams)

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
	//quotation_products_optionsModel := models.QuotationProductsOptions{}
	c.Bind(quotation_products_optionsModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newQuotationProductsOptions := models.NewQuotationProductsOptionsRepository(db.DB)
	err = newQuotationProductsOptions.Update(quotation_products_optionsModel)

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

	setFlashmessages(c, "success", "QuotationProductsOptions successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/quotation_products_options/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, quotation_products_options.QuotationProductsOptionsIndex(
		fmt.Sprintf("| Edit QuotationProductsOptions #%d", quotation_products_optionsModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		quotation_products_options.UpdateQuotationProductsOptions(*quotation_products_optionsModel, tz.(string)),
	))

}

func deleteQuotationProductsOptionsHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.QuotationProductsOptions{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newQuotationProductsOptions := models.NewQuotationProductsOptionsRepository(db.DB)
	err = newQuotationProductsOptions.Delete(idParams)
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

	setFlashmessages(c, "success", "QuotationProductsOptions successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/quotation_products_options/list")
}
