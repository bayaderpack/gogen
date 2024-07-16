package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/order_totals"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createOrderTotalsHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		order_totalsModel := models.OrderTotals{}
		c.Bind(&order_totalsModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newOrderTotals := models.NewOrderTotalsRepository(db.DB)
		newOrderTotals.Create(&order_totalsModel)

		setFlashmessages(c, "success", "OrderTotals created successfully!!")
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
	return renderView(c, order_totals.OrderTotalsIndex(
		"| Create OrderTotals",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		order_totals.CreateOrderTotals(),
	))

}

func order_totalsListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newOrderTotals := models.NewOrderTotalsRepository(db.DB)
	allOrderTotals, err := newOrderTotals.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's OrderTotals List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, order_totals.OrderTotalsIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		order_totals.OrderTotalsList(titlePage, allOrderTotals),
	))
}

func updateOrderTotalsHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newOrderTotals := models.NewOrderTotalsRepository(db.DB)

	order_totalsModel, err := newOrderTotals.GetSingle(idParams)

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
	//order_totalsModel := models.OrderTotals{}
	c.Bind(order_totalsModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newOrderTotals := models.NewOrderTotalsRepository(db.DB)
	err = newOrderTotals.Update(order_totalsModel)

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

	setFlashmessages(c, "success", "OrderTotals successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/order_totals/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, order_totals.OrderTotalsIndex(
		fmt.Sprintf("| Edit OrderTotals #%d", order_totalsModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		order_totals.UpdateOrderTotals(*order_totalsModel, tz.(string)),
	))

}

func deleteOrderTotalsHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.OrderTotals{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newOrderTotals := models.NewOrderTotalsRepository(db.DB)
	err = newOrderTotals.Delete(idParams)
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

	setFlashmessages(c, "success", "OrderTotals successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/order_totals/list")
}
