package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/order_products_options"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createOrderProductsOptionsHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		order_products_optionsModel := models.OrderProductsOptions{}
		c.Bind(&order_products_optionsModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newOrderProductsOptions := models.NewOrderProductsOptionsRepository(db.DB)
		newOrderProductsOptions.Create(&order_products_optionsModel)

		setFlashmessages(c, "success", "OrderProductsOptions created successfully!!")
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
	return renderView(c, order_products_options.OrderProductsOptionsIndex(
		"| Create OrderProductsOptions",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		order_products_options.CreateOrderProductsOptions(),
	))

}

func order_products_optionsListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newOrderProductsOptions := models.NewOrderProductsOptionsRepository(db.DB)
	allOrderProductsOptions, err := newOrderProductsOptions.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's OrderProductsOptions List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, order_products_options.OrderProductsOptionsIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		order_products_options.OrderProductsOptionsList(titlePage, allOrderProductsOptions),
	))
}

func updateOrderProductsOptionsHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newOrderProductsOptions := models.NewOrderProductsOptionsRepository(db.DB)

	order_products_optionsModel, err := newOrderProductsOptions.GetSingle(idParams)

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
	//order_products_optionsModel := models.OrderProductsOptions{}
	c.Bind(order_products_optionsModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newOrderProductsOptions := models.NewOrderProductsOptionsRepository(db.DB)
	err = newOrderProductsOptions.Update(order_products_optionsModel)

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

	setFlashmessages(c, "success", "OrderProductsOptions successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/order_products_options/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, order_products_options.OrderProductsOptionsIndex(
		fmt.Sprintf("| Edit OrderProductsOptions #%d", order_products_optionsModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		order_products_options.UpdateOrderProductsOptions(*order_products_optionsModel, tz.(string)),
	))

}

func deleteOrderProductsOptionsHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.OrderProductsOptions{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newOrderProductsOptions := models.NewOrderProductsOptionsRepository(db.DB)
	err = newOrderProductsOptions.Delete(idParams)
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

	setFlashmessages(c, "success", "OrderProductsOptions successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/order_products_options/list")
}
