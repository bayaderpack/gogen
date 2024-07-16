package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/quotation_products"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createQuotationProductsHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		quotation_productsModel := models.QuotationProducts{}
		c.Bind(&quotation_productsModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newQuotationProducts := models.NewQuotationProductsRepository(db.DB)
		newQuotationProducts.Create(&quotation_productsModel)

		setFlashmessages(c, "success", "QuotationProducts created successfully!!")
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
	return renderView(c, quotation_products.QuotationProductsIndex(
		"| Create QuotationProducts",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		quotation_products.CreateQuotationProducts(),
	))

}

func quotation_productsListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newQuotationProducts := models.NewQuotationProductsRepository(db.DB)
	allQuotationProducts, err := newQuotationProducts.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's QuotationProducts List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, quotation_products.QuotationProductsIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		quotation_products.QuotationProductsList(titlePage, allQuotationProducts),
	))
}

func UpdateQuotationProductsHandler(c *gin.Context) {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}
	newQuotationProducts := models.NewQuotationProductsRepository(db.DB)

	quotation_productsModel, err := newQuotationProducts.GetSingle(idParams)

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
	if c.Request.Method == "POST" {
		//quotation_productsModel := models.QuotationProducts{}
		c.Bind(quotation_productsModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		//newQuotationProducts := models.NewQuotationProductsRepository(db.DB)
		err = newQuotationProducts.Update(quotation_productsModel)

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

		setFlashmessages(c, "success", "QuotationProducts successfully updated!!")

		//	return c.Redirect(http.StatusSeeOther, "/quotation_products/list")
		//}
	} else {

		// username, _ := c.Get(username_key)
		// tz, _ := c.Get(tzone_key)
		Render(c, quotation_products.QuotationProductsIndex(
			fmt.Sprintf("| Edit QuotationProducts #%d", quotation_productsModel.QuotationProductID),
			"username.(string)",
			fromProtected,
			isError,
			getFlashmessages(c, "error"),
			getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
			quotation_products.UpdateQuotationProducts(*quotation_productsModel, "tz.(string)"),
		))
	}

}

func deleteQuotationProductsHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.QuotationProducts{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newQuotationProducts := models.NewQuotationProductsRepository(db.DB)
	err = newQuotationProducts.Delete(idParams)
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

	setFlashmessages(c, "success", "QuotationProducts successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/quotation_products/list")
}
