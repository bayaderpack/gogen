package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/order_products"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createOrderProductsHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		order_productsModel := models.OrderProducts{}
		c.Bind(&order_productsModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newOrderProducts := models.NewOrderProductsRepository(db.DB)
		newOrderProducts.Create(&order_productsModel)

		setFlashmessages(c, "success", "OrderProducts created successfully!!")
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
	return renderView(c, order_products.OrderProductsIndex(
		"| Create OrderProducts",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		order_products.CreateOrderProducts(),
	))

}

func order_productsListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newOrderProducts := models.NewOrderProductsRepository(db.DB)
	allOrderProducts, err := newOrderProducts.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's OrderProducts List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, order_products.OrderProductsIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		order_products.OrderProductsList(titlePage, allOrderProducts),
	))
}

func updateOrderProductsHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newOrderProducts := models.NewOrderProductsRepository(db.DB)

	order_productsModel, err := newOrderProducts.GetSingle(idParams)

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
	//order_productsModel := models.OrderProducts{}
	c.Bind(order_productsModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newOrderProducts := models.NewOrderProductsRepository(db.DB)
	err = newOrderProducts.Update(order_productsModel)

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

	setFlashmessages(c, "success", "OrderProducts successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/order_products/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, order_products.OrderProductsIndex(
		fmt.Sprintf("| Edit OrderProducts #%d", order_productsModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		order_products.UpdateOrderProducts(*order_productsModel, tz.(string)),
	))

}

func deleteOrderProductsHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.OrderProducts{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newOrderProducts := models.NewOrderProductsRepository(db.DB)
	err = newOrderProducts.Delete(idParams)
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

	setFlashmessages(c, "success", "OrderProducts successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/order_products/list")
}
