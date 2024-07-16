package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/payment_method"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createPaymentMethodHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		payment_methodModel := models.PaymentMethod{}
		c.Bind(&payment_methodModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newPaymentMethod := models.NewPaymentMethodRepository(db.DB)
		newPaymentMethod.Create(&payment_methodModel)

		setFlashmessages(c, "success", "PaymentMethod created successfully!!")
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
	return renderView(c, payment_method.PaymentMethodIndex(
		"| Create PaymentMethod",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		payment_method.CreatePaymentMethod(),
	))

}

func payment_methodListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newPaymentMethod := models.NewPaymentMethodRepository(db.DB)
	allPaymentMethod, err := newPaymentMethod.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's PaymentMethod List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, payment_method.PaymentMethodIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		payment_method.PaymentMethodList(titlePage, allPaymentMethod),
	))
}

func updatePaymentMethodHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newPaymentMethod := models.NewPaymentMethodRepository(db.DB)

	payment_methodModel, err := newPaymentMethod.GetSingle(idParams)

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
	//payment_methodModel := models.PaymentMethod{}
	c.Bind(payment_methodModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newPaymentMethod := models.NewPaymentMethodRepository(db.DB)
	err = newPaymentMethod.Update(payment_methodModel)

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

	setFlashmessages(c, "success", "PaymentMethod successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/payment_method/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, payment_method.PaymentMethodIndex(
		fmt.Sprintf("| Edit PaymentMethod #%d", payment_methodModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		payment_method.UpdatePaymentMethod(*payment_methodModel, tz.(string)),
	))

}

func deletePaymentMethodHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.PaymentMethod{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newPaymentMethod := models.NewPaymentMethodRepository(db.DB)
	err = newPaymentMethod.Delete(idParams)
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

	setFlashmessages(c, "success", "PaymentMethod successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/payment_method/list")
}
