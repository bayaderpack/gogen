package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/sample_payments"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createSamplePaymentsHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	sample_paymentsModel := models.SamplePayments{}
	c.Bind(&sample_paymentsModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newSamplePayments := models.NewSamplePaymentsRepository(db.DB)
	newSamplePayments.Create(&sample_paymentsModel)

setFlashmessages(c, "success", "SamplePayments created successfully!!")
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
	return renderView(c, sample_payments.SamplePaymentsIndex(
		"| Create SamplePayments",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		sample_payments.CreateSamplePayments(),
	))

}

func sample_paymentsListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newSamplePayments := models.NewSamplePaymentsRepository(db.DB)
	allSamplePayments, err := newSamplePayments.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's SamplePayments List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, sample_payments.SamplePaymentsIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		sample_payments.SamplePaymentsList(titlePage, allSamplePayments),
	))
}

func updateSamplePaymentsHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newSamplePayments := models.NewSamplePaymentsRepository(db.DB)
	
	sample_paymentsModel , err := newSamplePayments.GetSingle(idParams)

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
		//sample_paymentsModel := models.SamplePayments{}
	c.Bind(sample_paymentsModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newSamplePayments := models.NewSamplePaymentsRepository(db.DB)
	err = newSamplePayments.Update(sample_paymentsModel)

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


		setFlashmessages(c, "success", "SamplePayments successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/sample_payments/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, sample_payments.SamplePaymentsIndex(
		fmt.Sprintf("| Edit SamplePayments #%d", sample_paymentsModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		sample_payments.UpdateSamplePayments(*sample_paymentsModel, tz.(string)),
	))

}

func deleteSamplePaymentsHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.SamplePayments{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newSamplePayments := models.NewSamplePaymentsRepository(db.DB)
	err = newSamplePayments.Delete(idParams)
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

	setFlashmessages(c, "success", "SamplePayments successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/sample_payments/list")
}
