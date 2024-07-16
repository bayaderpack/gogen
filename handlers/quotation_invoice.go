package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/quotation_invoice"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createQuotationInvoiceHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		quotation_invoiceModel := models.QuotationInvoice{}
		c.Bind(&quotation_invoiceModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newQuotationInvoice := models.NewQuotationInvoiceRepository(db.DB)
		newQuotationInvoice.Create(&quotation_invoiceModel)

		setFlashmessages(c, "success", "QuotationInvoice created successfully!!")
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
	return renderView(c, quotation_invoice.QuotationInvoiceIndex(
		"| Create QuotationInvoice",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		quotation_invoice.CreateQuotationInvoice(),
	))

}

func quotation_invoiceListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newQuotationInvoice := models.NewQuotationInvoiceRepository(db.DB)
	allQuotationInvoice, err := newQuotationInvoice.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's QuotationInvoice List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, quotation_invoice.QuotationInvoiceIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		quotation_invoice.QuotationInvoiceList(titlePage, allQuotationInvoice),
	))
}

func updateQuotationInvoiceHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newQuotationInvoice := models.NewQuotationInvoiceRepository(db.DB)

	quotation_invoiceModel, err := newQuotationInvoice.GetSingle(idParams)

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
	//quotation_invoiceModel := models.QuotationInvoice{}
	c.Bind(quotation_invoiceModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newQuotationInvoice := models.NewQuotationInvoiceRepository(db.DB)
	err = newQuotationInvoice.Update(quotation_invoiceModel)

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

	setFlashmessages(c, "success", "QuotationInvoice successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/quotation_invoice/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, quotation_invoice.QuotationInvoiceIndex(
		fmt.Sprintf("| Edit QuotationInvoice #%d", quotation_invoiceModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		quotation_invoice.UpdateQuotationInvoice(*quotation_invoiceModel, tz.(string)),
	))

}

func deleteQuotationInvoiceHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.QuotationInvoice{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newQuotationInvoice := models.NewQuotationInvoiceRepository(db.DB)
	err = newQuotationInvoice.Delete(idParams)
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

	setFlashmessages(c, "success", "QuotationInvoice successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/quotation_invoice/list")
}
